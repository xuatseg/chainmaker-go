/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package verifier

import (
	"encoding/hex"
	"fmt"
	"sync"

	"chainmaker.org/chainmaker-go/module/consensus"
	"chainmaker.org/chainmaker-go/module/core/common"
	"chainmaker.org/chainmaker-go/module/core/common/coinbasemgr"
	"chainmaker.org/chainmaker-go/module/core/common/scheduler"
	"chainmaker.org/chainmaker-go/module/core/provider/conf"
	commonErrors "chainmaker.org/chainmaker/common/v2/errors"
	"chainmaker.org/chainmaker/common/v2/monitor"
	"chainmaker.org/chainmaker/common/v2/msgbus"
	"chainmaker.org/chainmaker/localconf/v2"
	commonpb "chainmaker.org/chainmaker/pb-go/v2/common"
	chainConfConfig "chainmaker.org/chainmaker/pb-go/v2/config"
	consensuspb "chainmaker.org/chainmaker/pb-go/v2/consensus"
	"chainmaker.org/chainmaker/protocol/v2"
	batch "chainmaker.org/chainmaker/txpool-batch/v2"
	"chainmaker.org/chainmaker/utils/v2"

	"github.com/gogo/protobuf/proto"
	"github.com/prometheus/client_golang/prometheus"
)

// BlockVerifierImpl implements BlockVerifier interface.
// Verify block and transactions.
//nolint: structcheck,unused

var ModuleNameCore = "Core"

type BlockVerifierImpl struct {
	chainId         string                   // chain id, to identity this chain
	msgBus          msgbus.MessageBus        // message bus
	txScheduler     protocol.TxScheduler     // scheduler orders tx batch into DAG form and returns a block
	snapshotManager protocol.SnapshotManager // snapshot manager
	ledgerCache     protocol.LedgerCache     // ledger cache
	blockchainStore protocol.BlockchainStore // blockchain store

	reentrantLocks *common.ReentrantLocks         // reentrant lock for avoid concurrent verify block
	proposalCache  protocol.ProposalCache         // proposal cache
	chainConf      protocol.ChainConf             // chain config
	ac             protocol.AccessControlProvider // access control manager
	log            protocol.Logger                // logger
	txPool         protocol.TxPool                // tx pool to check if tx is duplicate
	mu             sync.Mutex                     // to avoid concurrent map modify
	verifierBlock  *common.VerifierBlock
	storeHelper    conf.StoreHelper

	metricBlockVerifyTime *prometheus.HistogramVec // metrics monitor
	netService            protocol.NetService
	txFilter              protocol.TxFilter // Verify the transaction rules with TxFilter
}

type BlockVerifierConfig struct {
	ChainId         string
	MsgBus          msgbus.MessageBus
	SnapshotManager protocol.SnapshotManager
	BlockchainStore protocol.BlockchainStore
	LedgerCache     protocol.LedgerCache
	TxScheduler     protocol.TxScheduler
	ProposedCache   protocol.ProposalCache
	ChainConf       protocol.ChainConf
	AC              protocol.AccessControlProvider
	TxPool          protocol.TxPool
	TxFilter        protocol.TxFilter
	VmMgr           protocol.VmManager
	StoreHelper     conf.StoreHelper
	NetService      protocol.NetService
}

func NewBlockVerifier(config BlockVerifierConfig, log protocol.Logger) (protocol.BlockVerifier, error) {
	v := &BlockVerifierImpl{
		chainId:         config.ChainId,
		msgBus:          config.MsgBus,
		txScheduler:     config.TxScheduler,
		snapshotManager: config.SnapshotManager,
		ledgerCache:     config.LedgerCache,
		blockchainStore: config.BlockchainStore,
		reentrantLocks: &common.ReentrantLocks{
			ReentrantLocks: make(map[string]interface{}),
		},
		proposalCache: config.ProposedCache,
		chainConf:     config.ChainConf,
		ac:            config.AC,
		log:           log,
		txPool:        config.TxPool,
		storeHelper:   config.StoreHelper,
		netService:    config.NetService,
		txFilter:      config.TxFilter,
	}

	verifyConf := &common.VerifierBlockConf{
		ChainConf:       v.chainConf,
		Log:             v.log,
		LedgerCache:     v.ledgerCache,
		Ac:              v.ac,
		SnapshotManager: v.snapshotManager,
		TxPool:          v.txPool,
		BlockchainStore: v.blockchainStore,
		ProposalCache:   v.proposalCache,
		VmMgr:           config.VmMgr,
		StoreHelper:     config.StoreHelper,
		TxScheduler:     config.TxScheduler,
		TxFilter:        config.TxFilter,
	}
	v.verifierBlock = common.NewVerifierBlock(verifyConf)

	if localconf.ChainMakerConfig.MonitorConfig.Enabled {
		v.metricBlockVerifyTime = monitor.NewHistogramVec(monitor.SUBSYSTEM_CORE_VERIFIER, "metric_block_verify_time",
			"block verify time metric", []float64{0.005, 0.01, 0.015, 0.05, 0.1, 1, 2, 5, 10}, "chainId")
	}

	// v220_compat Deprecated
	config.ChainConf.AddWatch(v) //nolint: staticcheck
	config.MsgBus.Register(msgbus.ChainConfig, v)

	return v, nil
}

// OnMessage contract event data is a []string, hexToString(proto.Marshal(data))
func (v *BlockVerifierImpl) OnMessage(msg *msgbus.Message) {
	switch msg.Topic {
	case msgbus.ChainConfig:
		dataStr, ok := msg.Payload.([]string)
		if !ok {
			return
		}
		dataBytes, err := hex.DecodeString(dataStr[0])
		if err != nil {
			v.log.Warn(err)
			return
		}
		chainConfig := &chainConfConfig.ChainConfig{}
		err = proto.Unmarshal(dataBytes, chainConfig)
		if err != nil {
			v.log.Warn(err)
			return
		}
		v.chainConf.ChainConfig().Block = chainConfig.Block
		protocol.ParametersValueMaxLength = chainConfig.Block.TxParameterSize * 1024 * 1024
		if chainConfig.Block.TxParameterSize <= 0 {
			protocol.ParametersValueMaxLength = protocol.DefaultParametersValueMaxSize * 1024 * 1024
		}
		v.log.Infof("[BlockVerifierImpl] receive msg, topic: %s, blockverify[%v]",
			msg.Topic.String(), v.chainConf.ChainConfig().Block)
	default:

	}
}

func (v *BlockVerifierImpl) OnQuit() {
	// nothing, implement Subscriber interface
}

// VerifyBlock to check if block is valid
func (v *BlockVerifierImpl) VerifyBlock(block *commonpb.Block, mode protocol.VerifyMode) (err error) {
	_, err = v.verifyBlock(block, mode)
	return err
}

func (v *BlockVerifierImpl) VerifyBlockSync(block *commonpb.Block,
	mode protocol.VerifyMode) (result *consensuspb.VerifyResult, err error) {
	return v.verifyBlock(block, mode)
}

// VerifyBlock to check if block is valid
func (v *BlockVerifierImpl) verifyBlock(block *commonpb.Block, mode protocol.VerifyMode) (
	result *consensuspb.VerifyResult, err error) {

	blockVersion := block.Header.BlockVersion
	startTick := utils.CurrentTimeMillisSeconds()
	if err = utils.IsEmptyBlock(block); err != nil {
		v.log.Error(err)
		return nil, err
	}

	v.log.Debugf("verify receive [%d](%x,%d,%d), from sync %d",
		block.Header.BlockHeight, block.Header.BlockHash, block.Header.TxCount, len(block.Txs), mode)
	// avoid concurrent verify, only one block hash can be verified at the same time
	if !v.reentrantLocks.Lock(string(block.Header.BlockHash)) {
		v.log.Warnf("block(%d,%x) concurrent verify, yield", block.Header.BlockHeight, block.Header.BlockHash)
		return nil, commonErrors.ErrConcurrentVerify
	}
	defer v.reentrantLocks.Unlock(string(block.Header.BlockHash))

	// No duplicate verify
	result, isRepeat := v.verifyRepeat(block, startTick, mode)
	if isRepeat {
		return result, nil
	}
	var contractEventMap map[string][]*commonpb.ContractEvent

	// avoid to recover the committed block.
	lastBlock, err := v.verifierBlock.FetchLastBlock(block)
	if err != nil {
		return nil, err
	}

	startPoolTick := utils.CurrentTimeMillisSeconds()
	newBlock, batchIds, err := common.RecoverBlock(block, mode, v.chainConf, v.txPool, v.ac, v.netService, v.log)
	if err != nil {
		return nil, err
	}
	lastPool := utils.CurrentTimeMillisSeconds() - startPoolTick

	txRWSetMap, contractEventMap, timeLasts, rwSetVerifyFailTx, err := v.validateBlock(newBlock, lastBlock, mode)

	verifyResult := &consensuspb.VerifyResult{}
	if err != nil {
		v.log.Warnf("verify failed [%d](%x),preBlockHash:%x, %s",
			newBlock.Header.BlockHeight, newBlock.Header.BlockHash, newBlock.Header.PreBlockHash, err.Error())

		// if mode equal consensus verify, publish to consensus verify result signal
		if protocol.CONSENSUS_VERIFY == mode {
			v.log.DebugDynamic(func() string {
				return fmt.Sprintf("publish verfiy failed rw set txs, "+
					"block height:%d, err: %s", newBlock.Header.BlockHeight, err.Error())
			})
			verifyResult = parseVerifyResult(newBlock, false, txRWSetMap, rwSetVerifyFailTx)
			v.msgBus.Publish(msgbus.VerifyResult, verifyResult)
		}

		// rollback sql
		if sqlErr := v.storeHelper.RollBack(newBlock, v.blockchainStore); sqlErr != nil {
			v.log.Errorf("block [%d] rollback sql failed: %s", newBlock.Header.BlockHeight, sqlErr)
		}
		return verifyResult, err
	}

	snapshot := v.snapshotManager.GetSnapshot(lastBlock, newBlock)
	if coinbasemgr.IsOptimizeChargeGasEnabled(v.chainConf) {
		if err = scheduler.VerifyOptimizeChargeGasTx(newBlock, snapshot, v.ac, blockVersion); err != nil {
			v.log.Warnf("verify failed [%d](%x), %s",
				newBlock.Header.BlockHeight, newBlock.Header.BlockHash, err.Error())
			verifyResult = parseVerifyResult(newBlock, true, txRWSetMap, rwSetVerifyFailTx)
			if protocol.CONSENSUS_VERIFY == mode {
				v.msgBus.Publish(msgbus.VerifyResult, verifyResult)
			}
			return verifyResult, err
		}
	}

	// sync mode, need to verify consensus vote signature
	beginConsensCheck := utils.CurrentTimeMillisSeconds()
	if protocol.SYNC_VERIFY == mode {
		if err = v.verifyVoteSig(newBlock); err != nil {
			v.log.Warnf("verify failed [%d](%x), votesig %s",
				newBlock.Header.BlockHeight, newBlock.Header.BlockHash, err.Error())
			return nil, err
		}
	}
	consensusCheckUsed := utils.CurrentTimeMillisSeconds() - beginConsensCheck

	// verify success, cache block and read write set
	// solo need this，too！！！
	v.log.Debugf("set proposed block(%d,%x)", newBlock.Header.BlockHeight, newBlock.Header.BlockHash)
	if err = v.proposalCache.SetProposedBlock(newBlock, txRWSetMap, contractEventMap, false); err != nil {
		return nil, err
	}

	// mark transactions in block as pending status in txpool
	if common.TxPoolType == batch.TxPoolType {
		v.txPool.AddTxBatchesToPendingCache(batchIds, newBlock.Header.BlockHeight)
	} else {
		v.txPool.AddTxsToPendingCache(coinbasemgr.FilterCoinBaseTxOrGasTx(newBlock.Txs), newBlock.Header.BlockHeight)
	}

	// if mode equal consensus verify, publish to consensus verify result signal
	if protocol.CONSENSUS_VERIFY == mode {
		verifyResult = parseVerifyResult(newBlock, true, txRWSetMap, nil)
		v.msgBus.Publish(msgbus.VerifyResult, verifyResult)
	}
	elapsed := utils.CurrentTimeMillisSeconds() - startTick
	v.log.Infof("verify success [%d,%x]"+
		"(blockSig:%d,vm:%d,txVerify:%d,txRoot:%d,pool:%d,consensusCheckUsed:%d,total:%d)",
		newBlock.Header.BlockHeight, newBlock.Header.BlockHash, timeLasts[common.BlockSig], timeLasts[common.VM],
		timeLasts[common.TxVerify], timeLasts[common.TxRoot], lastPool, consensusCheckUsed, elapsed)

	// monitor config open case
	if localconf.ChainMakerConfig.MonitorConfig.Enabled {
		v.metricBlockVerifyTime.WithLabelValues(v.chainId).Observe(float64(elapsed) / 1000)
	}

	return verifyResult, nil
}

// VerifyBlockWithRwSets to check if block is valid
func (v *BlockVerifierImpl) VerifyBlockWithRwSets(block *commonpb.Block,
	rwsets []*commonpb.TxRWSet, mode protocol.VerifyMode) (err error) {

	startTick := utils.CurrentTimeMillisSeconds()
	if err = utils.IsEmptyBlock(block); err != nil {
		v.log.Error(err)
		v.log.Debugf("empty block. height:%+v, hash:%+v, chainId:%+v, preHash:%+v, signature:%+v",
			block.Header.BlockHeight, block.Header.BlockHash,
			block.Header.ChainId, block.Header.PreBlockHash, block.Header.Signature)
		return err
	}

	v.log.Debugf("VerifyBlockWithRwSets receive [%d](%x,%d,%d), from sync %d",
		block.Header.BlockHeight, block.Header.BlockHash, block.Header.TxCount, len(block.Txs), mode)
	// avoid concurrent verify, only one block hash can be verified at the same time
	if !v.reentrantLocks.Lock(string(block.Header.BlockHash)) {
		v.log.Warnf("block(%d,%x) concurrent verify, yield", block.Header.BlockHeight, block.Header.BlockHash)
		return commonErrors.ErrConcurrentVerify
	}
	defer v.reentrantLocks.Unlock(string(block.Header.BlockHash))
	// No duplicate verify
	_, isRepeat := v.verifyRepeat(block, startTick, mode)
	if isRepeat {
		return nil
	}

	txRWSetMap := make(map[string]*commonpb.TxRWSet)
	for _, txRWSet := range rwsets {
		if txRWSet != nil {
			txRWSetMap[txRWSet.TxId] = txRWSet
		}
	}

	// avoid to recover the committed block.
	lastBlock, err := v.verifierBlock.FetchLastBlock(block)
	if err != nil {
		return err
	}

	startPoolTick := utils.CurrentTimeMillisSeconds()
	newBlock, batchIds, err := common.RecoverBlock(block, mode, v.chainConf, v.txPool, v.ac, v.netService, v.log)
	if err != nil {
		return err
	}
	lastPool := utils.CurrentTimeMillisSeconds() - startPoolTick
	contractEventMap, timeLasts, err := v.validateBlockWithRWSets(newBlock, lastBlock, mode, txRWSetMap)
	if err != nil {
		v.log.Warnf("verify failed [%d](%x),preBlockHash:%x, %s",
			newBlock.Header.BlockHeight, newBlock.Header.BlockHash, newBlock.Header.PreBlockHash, err.Error())

		// rollback sql
		if sqlErr := v.storeHelper.RollBack(newBlock, v.blockchainStore); sqlErr != nil {
			v.log.Errorf("block [%d] rollback sql failed: %s", newBlock.Header.BlockHeight, sqlErr)
		}
		return err
	}

	// sync mode, need to verify consensus vote signature
	beginConsensCheck := utils.CurrentTimeMillisSeconds()
	if protocol.SYNC_VERIFY == mode {
		if err = v.verifyVoteSig(newBlock); err != nil {
			v.log.Warnf("verify failed [%d](%x), votesig %s",
				newBlock.Header.BlockHeight, newBlock.Header.BlockHash, err.Error())
			return err
		}
	}
	consensusCheckUsed := utils.CurrentTimeMillisSeconds() - beginConsensCheck

	// verify success, cache block and read write set
	// solo need this，too！！！
	v.log.Debugf("set proposed block(%d,%x)", newBlock.Header.BlockHeight, newBlock.Header.BlockHash)
	if err = v.proposalCache.SetProposedBlock(newBlock, txRWSetMap, contractEventMap, false); err != nil {
		return err
	}
	currSnapshot := v.snapshotManager.NewSnapshot(lastBlock, block)
	currSnapshot.ApplyBlock(block, txRWSetMap)

	// mark transactions in block as pending status in txpool
	if common.TxPoolType == batch.TxPoolType {
		v.txPool.AddTxBatchesToPendingCache(batchIds, newBlock.Header.BlockHeight)
	} else {
		v.txPool.AddTxsToPendingCache(coinbasemgr.FilterCoinBaseTxOrGasTx(newBlock.Txs), newBlock.Header.BlockHeight)
	}

	// if mode equal consensus verify, publish to consensus verify result signal
	if protocol.CONSENSUS_VERIFY == mode {
		v.msgBus.Publish(msgbus.VerifyResult, parseVerifyResult(newBlock, true, txRWSetMap, nil))
	}
	elapsed := utils.CurrentTimeMillisSeconds() - startTick
	v.log.Infof("verify success [%d,%x]"+
		"(blockSig:%d,vm:%d,txVerify:%d,txRoot:%d,pool:%d,consensusCheckUsed:%d,total:%d)",
		newBlock.Header.BlockHeight, newBlock.Header.BlockHash, timeLasts[common.BlockSig], timeLasts[common.VM],
		timeLasts[common.TxVerify], timeLasts[common.TxRoot], lastPool, consensusCheckUsed, elapsed)

	// monitor config open case
	if localconf.ChainMakerConfig.MonitorConfig.Enabled {
		v.metricBlockVerifyTime.WithLabelValues(v.chainId).Observe(float64(elapsed) / 1000)
	}
	return nil
}

// Module return module name core
func (v *BlockVerifierImpl) Module() string {
	return ModuleNameCore
}

// Watch if tx paramer size less than or equal 0, set the parameter value max length by default parameter value max size
func (v *BlockVerifierImpl) Watch(chainConfig *chainConfConfig.ChainConfig) error {
	v.chainConf.ChainConfig().Block = chainConfig.Block
	protocol.ParametersValueMaxLength = chainConfig.Block.TxParameterSize * 1024 * 1024
	if chainConfig.Block.TxParameterSize <= 0 {
		protocol.ParametersValueMaxLength = protocol.DefaultParametersValueMaxSize * 1024 * 1024
	}
	v.log.Infof("update chainconf,blockverify[%v]", v.chainConf.ChainConfig().Block)
	return nil
}

// validateBlock validate block
func (v *BlockVerifierImpl) validateBlock(block,
	lastBlock *commonpb.Block, mode protocol.VerifyMode) (map[string]*commonpb.TxRWSet,
	map[string][]*commonpb.ContractEvent,
	map[string]int64, *common.RwSetVerifyFailTx, error) {
	hashType := v.chainConf.ChainConfig().Crypto.Hash
	timeLasts := make(map[string]int64)
	var err error
	var txCapacity uint32
	if coinbasemgr.IsOptimizeChargeGasEnabled(v.chainConf) {
		txCapacity = v.chainConf.ChainConfig().Block.BlockTxCapacity + 1
	} else {
		txCapacity = v.chainConf.ChainConfig().Block.BlockTxCapacity
	}

	if block.Header.TxCount > txCapacity {
		return nil, nil, timeLasts, nil, fmt.Errorf("txcapacity expect <= %d, got %d)", txCapacity, block.Header.TxCount)
	}

	if err = common.IsTxCountValid(block); err != nil {
		return nil, nil, timeLasts, nil, err
	}

	// proposed height == proposing height - 1
	proposedHeight := lastBlock.Header.BlockHeight
	// check if this block height is 1 bigger than last block height
	lastBlockHash := lastBlock.Header.BlockHash
	err = v.checkPreBlock_MAXBFT(block, lastBlock, err, lastBlockHash, proposedHeight)
	if err != nil {
		return nil, nil, timeLasts, nil, err
	}

	// ValidateBlock block by verifier
	return v.verifierBlock.ValidateBlock(block, lastBlock, hashType, timeLasts, mode)
}

// validateBlockWithRWSets validate block with rw sets
func (v *BlockVerifierImpl) validateBlockWithRWSets(block, lastBlock *commonpb.Block, mode protocol.VerifyMode,
	txRWSetMap map[string]*commonpb.TxRWSet) (
	map[string][]*commonpb.ContractEvent, map[string]int64, error) {
	hashType := v.chainConf.ChainConfig().Crypto.Hash
	timeLasts := make(map[string]int64)
	var err error
	// proposed height == proposing height - 1
	proposedHeight := lastBlock.Header.BlockHeight
	// check if this block height is 1 bigger than last block height
	lastBlockHash := lastBlock.Header.BlockHash
	err = common.CheckPreBlock(block, lastBlockHash, proposedHeight)

	if err != nil {
		return nil, timeLasts, err
	}

	// ValidateBlockWithRWSets block by verifier
	return v.verifierBlock.ValidateBlockWithRWSets(block, hashType, timeLasts, txRWSetMap, mode)
}

// checkPreBlock_MAXBFT check prepare block by max bft consensus type
func (v *BlockVerifierImpl) checkPreBlock_MAXBFT(block *commonpb.Block, lastBlock *commonpb.Block, err error,
	lastBlockHash []byte, proposedHeight uint64) error {

	if block.Header.BlockHeight == lastBlock.Header.BlockHeight+1 {
		if err = common.IsPreHashValid(block, lastBlock.Header.BlockHash); err != nil {
			return err
		}
	} else {
		// for maxbft consensus type
		proposedBlock, _ := v.proposalCache.GetProposedBlockByHashAndHeight(
			block.Header.PreBlockHash, block.Header.BlockHeight-1)
		if proposedBlock == nil {
			return fmt.Errorf(
				"no last block found [%d](%x) %s",
				block.Header.BlockHeight-1,
				block.Header.PreBlockHash,
				err,
			)
		}
	}

	return nil
}

// verifyVoteSig verify vote signatures
func (v *BlockVerifierImpl) verifyVoteSig(block *commonpb.Block) error {
	return consensus.VerifyBlockSignatures(v.chainConf, v.ac, v.blockchainStore, block, v.ledgerCache)
}

// parseVerifyResult pater verify result, return verify result
func parseVerifyResult(
	block *commonpb.Block,
	isValid bool,
	txsRwSet map[string]*commonpb.TxRWSet,
	rwSetVerifyFailTxs *common.RwSetVerifyFailTx) *consensuspb.VerifyResult {
	verifyResult := &consensuspb.VerifyResult{
		VerifiedBlock: block,
		TxsRwSet:      txsRwSet,
	}
	if isValid {
		verifyResult.Code = consensuspb.VerifyResult_SUCCESS
		verifyResult.Msg = "OK"
	} else {
		verifyResult.Msg = "FAIL"
		verifyResult.Code = consensuspb.VerifyResult_FAIL
		// if rw set verify fail tx not nil, set the verify result tx ids and block height
		if rwSetVerifyFailTxs != nil {
			verifyResult.RwSetVerifyFailTxs = &consensuspb.RwSetVerifyFailTxs{
				TxIds:       rwSetVerifyFailTxs.TxIds,
				BlockHeight: rwSetVerifyFailTxs.BlockHeight,
			}
		}
	}
	return verifyResult
}

// cutBlocks cut blocks
func (v *BlockVerifierImpl) cutBlocks(blocksToCut []*commonpb.Block, blockToKeep *commonpb.Block) {

	if common.TxPoolType == batch.TxPoolType {
		err := v.cutBlocksForBatchPool(blocksToCut, blockToKeep)
		if err != nil {
			v.log.Warnf(fmt.Sprintf("cut block[%d] failed, err:%v", blockToKeep.Header.BlockHeight, err))
		}
		return
	}

	cutTxs := make([]*commonpb.Transaction, 0)
	txMap := make(map[string]interface{})
	// make map, the key is tx id, and the value tx
	for _, tx := range blockToKeep.Txs {
		txMap[tx.Payload.TxId] = struct{}{}
	}
	// collect the tx map
	for _, blockToCut := range blocksToCut {
		v.log.Infof("cut block hash: %x, height: %v", blockToCut.Header.BlockHash, blockToCut.Header.BlockHeight)
		for _, txToCut := range blockToCut.Txs {
			if _, ok := txMap[txToCut.Payload.TxId]; ok {
				// this transaction is kept, do NOT cut it.
				continue
			}
			v.log.Debugf("cut tx hash: %s", txToCut.Payload.TxId)
			cutTxs = append(cutTxs, txToCut)
		}
	}
	// if cut txs not nil, retry txs to tx pool
	if len(cutTxs) > 0 {
		common.RetryAndRemoveTxs(v.txPool, cutTxs, nil, v.log)
	}
}

func (v *BlockVerifierImpl) cutBlocksForBatchPool(blocksToCut []*commonpb.Block, blockToKeep *commonpb.Block) error {

	keepBatchIdsMap := make(map[string]interface{})
	batchIds, _, err := common.GetBatchIds(blockToKeep)
	if err != nil {
		v.log.Errorf("get batch ids from keep block[%d,%x] failed, err:%v",
			blockToKeep.Header.BlockHeight, blockToKeep.Header.BlockHash, err)
		return err
	}

	for _, batchId := range batchIds {
		keepBatchIdsMap[batchId] = struct{}{}
	}

	finalCutBatchIds := make([]string, 0)
	for _, blockToCut := range blocksToCut {
		v.log.Infof("cut block hash: %x, height: %v", blockToCut.Header.BlockHash, blockToCut.Header.BlockHeight)
		cutBatchIds, _, err := common.GetBatchIds(blockToCut)
		if err != nil {
			v.log.Warnf("get batch ids from removed block[%d,%x] failed, err:%v",
				blockToCut.Header.BlockHeight, blockToCut.Header.BlockHash, err)
			continue
		}
		for _, cutBatchId := range cutBatchIds {
			if _, ok := keepBatchIdsMap[cutBatchId]; ok {
				// this transaction is kept, do NOT cut it.
				continue
			}
			v.log.Debugf("cut tx batchId: %s", cutBatchId)
			finalCutBatchIds = append(finalCutBatchIds, cutBatchId)
		}
	}

	if len(finalCutBatchIds) > 0 {
		v.txPool.RetryAndRemoveTxBatches(finalCutBatchIds, nil)
	}

	return nil
}

// verifyRepeat to check if the block has verified before
func (v *BlockVerifierImpl) verifyRepeat(block *commonpb.Block, startTick int64,
	mode protocol.VerifyMode) (result *consensuspb.VerifyResult, isRepeat bool) {
	b, txRwSet, _ := v.proposalCache.GetProposedBlock(block)
	if b == nil {
		return nil, false
	}
	if consensuspb.ConsensusType_SOLO != v.chainConf.ChainConfig().Consensus.Type ||
		v.chainConf.ChainConfig().Contract.EnableSqlSupport {
		elapsed := utils.CurrentTimeMillisSeconds() - startTick
		// the block has verified before
		v.log.Infof("verify success repeat [%d](%x), total: %d", b.Header.BlockHeight, b.Header.BlockHash, elapsed)
		if protocol.CONSENSUS_VERIFY == mode {
			// consensus mode, publish verify result to message bus
			result = parseVerifyResult(b, true, txRwSet, nil)
			v.msgBus.Publish(msgbus.VerifyResult, result)
		}
		lastBlock, _ := v.proposalCache.GetProposedBlockByHashAndHeight(
			b.Header.PreBlockHash, b.Header.BlockHeight-1)
		if lastBlock == nil {
			v.log.Debugf(
				"no pre-block be found, preHeight:%d, preBlockHash:%x",
				b.Header.BlockHeight-1,
				b.Header.PreBlockHash,
			)
			return result, true
		}
		cutBlocks := v.proposalCache.KeepProposedBlock(lastBlock.Header.BlockHash, lastBlock.Header.BlockHeight)
		if len(cutBlocks) > 0 {
			v.log.Infof(
				"received block hash: %s, height: %v",
				hex.EncodeToString(lastBlock.Header.BlockHash),
				lastBlock.Header.BlockHeight,
			)
			v.cutBlocks(cutBlocks, lastBlock)
		}

		return result, true
	}
	return nil, false
}
