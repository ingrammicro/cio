// Copyright (c) 2017-2021 Ingram Micro Inc.

package network

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewListenerServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewListenerService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListListeners(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	listenersIn := testdata.GetListenerData()
	ListListenersMocked(t, loadBalancersIn[0].ID, listenersIn)
	ListListenersFailErrMocked(t, loadBalancersIn[0].ID, listenersIn)
	ListListenersFailStatusMocked(t, loadBalancersIn[0].ID, listenersIn)
	ListListenersFailJSONMocked(t, loadBalancersIn[0].ID, listenersIn)
}

func TestGetListener(t *testing.T) {
	listenersIn := testdata.GetListenerData()
	for _, listenerIn := range listenersIn {
		GetListenerMocked(t, listenerIn)
		GetListenerFailErrMocked(t, listenerIn)
		GetListenerFailStatusMocked(t, listenerIn)
		GetListenerFailJSONMocked(t, listenerIn)
	}
}

func TestCreateListener(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	listenersIn := testdata.GetListenerData()
	for _, listenerIn := range listenersIn {
		CreateListenerMocked(t, loadBalancersIn[0].ID, listenerIn)
		CreateListenerFailErrMocked(t, loadBalancersIn[0].ID, listenerIn)
		CreateListenerFailStatusMocked(t, loadBalancersIn[0].ID, listenerIn)
		CreateListenerFailJSONMocked(t, loadBalancersIn[0].ID, listenerIn)
	}
}

func TestUpdateListener(t *testing.T) {
	listenersIn := testdata.GetListenerData()
	for _, listenerIn := range listenersIn {
		UpdateListenerMocked(t, listenerIn)
		UpdateListenerFailErrMocked(t, listenerIn)
		UpdateListenerFailStatusMocked(t, listenerIn)
		UpdateListenerFailJSONMocked(t, listenerIn)
	}
}

func TestDeleteListener(t *testing.T) {
	listenersIn := testdata.GetListenerData()
	for _, listenerIn := range listenersIn {
		DeleteListenerMocked(t, listenerIn)
		DeleteListenerFailErrMocked(t, listenerIn)
		DeleteListenerFailStatusMocked(t, listenerIn)
		DeleteListenerFailJSONMocked(t, listenerIn)
	}
}

func TestRetryListener(t *testing.T) {
	listenersIn := testdata.GetListenerData()
	for _, listenerIn := range listenersIn {
		RetryListenerMocked(t, listenerIn)
		RetryListenerFailErrMocked(t, listenerIn)
		RetryListenerFailStatusMocked(t, listenerIn)
		RetryListenerFailJSONMocked(t, listenerIn)
	}
}

func TestListRules(t *testing.T) {
	listenersIn := testdata.GetListenerData()
	rulesIn := testdata.GetListenerRuleData()
	ListRulesMocked(t, listenersIn[0].ID, rulesIn)
	ListRulesFailErrMocked(t, listenersIn[0].ID, rulesIn)
	ListRulesFailStatusMocked(t, listenersIn[0].ID, rulesIn)
	ListRulesFailJSONMocked(t, listenersIn[0].ID, rulesIn)
}

func TestCreateRule(t *testing.T) {
	listenersIn := testdata.GetListenerData()
	rulesIn := testdata.GetListenerRuleData()
	for _, ruleIn := range rulesIn {
		CreateRuleMocked(t, listenersIn[0].ID, ruleIn)
		CreateRuleFailErrMocked(t, listenersIn[0].ID, ruleIn)
		CreateRuleFailStatusMocked(t, listenersIn[0].ID, ruleIn)
		CreateRuleFailJSONMocked(t, listenersIn[0].ID, ruleIn)
	}
}

func TestUpdateRule(t *testing.T) {
	listenersIn := testdata.GetListenerData()
	rulesIn := testdata.GetListenerRuleData()
	for _, ruleIn := range rulesIn {
		UpdateRuleMocked(t, listenersIn[0].ID, ruleIn)
		UpdateRuleFailErrMocked(t, listenersIn[0].ID, ruleIn)
		UpdateRuleFailStatusMocked(t, listenersIn[0].ID, ruleIn)
		UpdateRuleFailJSONMocked(t, listenersIn[0].ID, ruleIn)
	}
}

func TestDeleteRule(t *testing.T) {
	listenersIn := testdata.GetListenerData()
	rulesIn := testdata.GetListenerRuleData()
	for _, ruleIn := range rulesIn {
		DeleteRuleMocked(t, listenersIn[0].ID, ruleIn)
		DeleteRuleFailErrMocked(t, listenersIn[0].ID, ruleIn)
		DeleteRuleFailStatusMocked(t, listenersIn[0].ID, ruleIn)
	}
}
