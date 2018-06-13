// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import chain "github.com/kyokan/plasma/chain"

import mock "github.com/stretchr/testify/mock"
import util "github.com/kyokan/plasma/util"

// InvalidBlockDao is an autogenerated mock type for the InvalidBlockDao type
type InvalidBlockDao struct {
	mock.Mock
}

// Get provides a mock function with given fields: blkHash
func (_m *InvalidBlockDao) Get(blkHash util.Hash) (*chain.Block, error) {
	ret := _m.Called(blkHash)

	var r0 *chain.Block
	if rf, ok := ret.Get(0).(func(util.Hash) *chain.Block); ok {
		r0 = rf(blkHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chain.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(util.Hash) error); ok {
		r1 = rf(blkHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: blk
func (_m *InvalidBlockDao) Save(blk *chain.Block) error {
	ret := _m.Called(blk)

	var r0 error
	if rf, ok := ret.Get(0).(func(*chain.Block) error); ok {
		r0 = rf(blk)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}