// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/lasthyphen/beacongo/utils/units"
	"github.com/lasthyphen/beacongo/vms/platformvm/reward"
)

var (
	fujiGenesisConfigJSON = `{
		"networkID": 5,
		"allocations": [
			{
				"ethAddr": "0xd86b355443158939c2f1b2A00961F8453b33E74E",
				"djtxAddr": "X-dijets1v8wat5z4cxh7wh873d7n6d9m6mpnynr8sgl059",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 40000000000000000
					}
				]
			},
			{
				"ethAddr": "0x1B00f59fff05F6591c13e32740377eAE72661061",
				"djtxAddr": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"initialAmount": 32000000000000000,
				"unlockSchedule": []
			},
			{
				"ethAddr": "0x80231567cD6E270c8360B80b97034Ec26dad83b8",
				"djtxAddr": "X-dijets16yd4ams4xdp9c6ht9zfnp90225ukwmqnj964sw",
				"initialAmount": 32000000000000000,
				"unlockSchedule": []
			},
			{
				"ethAddr": "0x2CF25c0323F52bb8d3f46EBc2DFcd34Ba057b781",
				"djtxAddr": "X-dijets1mc6ar37ggvvh80ezkwnnyrkfummc32k0z9v9la",
				"initialAmount": 32000000000000000,
				"unlockSchedule": []
			},
			{
				"ethAddr": "0x1a8398064cbee749c62cc2943da234f75350390e",
				"djtxAddr": "X-dijets1u3c5j2f3kg38eg6ryflljj3n7ark48gtr7s6dh",
				"initialAmount": 32000000000000000,
				"unlockSchedule": []
			},
			{
				"ethAddr": "0x1a8398064cbee749c62cc2943da234f75350390e",
				"djtxAddr": "X-dijets1u3c5j2f3kg38eg6ryflljj3n7ark48gtr7s6dh",
				"initialAmount": 32000000000000000,
				"unlockSchedule": []
			},
			{
				"ethAddr": "0x1a8398064cbee749c62cc2943da234f75350390e",
				"djtxAddr": "X-dijets1u3c5j2f3kg38eg6ryflljj3n7ark48gtr7s6dh",
				"initialAmount": 32000000000000000,
				"unlockSchedule": []
			},
			{
				"ethAddr": "0x1a8398064cbee749c62cc2943da234f75350390e",
				"djtxAddr": "X-dijets1u3c5j2f3kg38eg6ryflljj3n7ark48gtr7s6dh",
				"initialAmount": 32000000000000000,
				"unlockSchedule": []
			},
			{
				"ethAddr": "0x1a8398064cbee749c62cc2943da234f75350390e",
				"djtxAddr": "X-dijets1u3c5j2f3kg38eg6ryflljj3n7ark48gtr7s6dh",
				"initialAmount": 32000000000000000,
				"unlockSchedule": []
			},
			{
				"ethAddr": "0x1a8398064cbee749c62cc2943da234f75350390e",
				"djtxAddr": "X-dijets1u3c5j2f3kg38eg6ryflljj3n7ark48gtr7s6dh",
				"initialAmount": 32000000000000000,
				"unlockSchedule": []
			},
			{
				"ethAddr": "0x1a8398064cbee749c62cc2943da234f75350390e",
				"djtxAddr": "X-dijets1u3c5j2f3kg38eg6ryflljj3n7ark48gtr7s6dh",
				"initialAmount": 32000000000000000,
				"unlockSchedule": []
			}
		],
		"startTime": 1599696000,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 54000,
		"initialStakedFunds": [
			"X-fuji1wycv8n7d2fg9aq6unp23pnj4q0arv03ysya8jw"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg",
				"rewardAddress": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"delegationFee": 1000000
			},
			{
				"nodeID": "NodeID-MFrZFVCXPv5iCn6M9K6XduxGTYp891xXZ",
				"rewardAddress": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"delegationFee": 500000
			},
			{
				"nodeID": "NodeID-NFBbbJ4qCmNaCzeW7sxErhvWqvEQMnYcN",
				"rewardAddress": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"delegationFee": 250000
			},
			{
				"nodeID": "NodeID-GWPcbFJZFfZreETSoWjPimr846mXEKCtu",
				"rewardAddress": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"delegationFee": 125000
			},
			{
				"nodeID": "NodeID-P7oB2McjBGgW2NXXWVYjV8JEDFoW9xDE5",
				"rewardAddress": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"delegationFee": 62500
			},
			{
				"nodeID": "NodeID-8Jx5tzgdZbGAmmXMkEfjbMPenEb2rCNVv",
				"rewardAddress": "X-dijets1dep9x84as7nku059x8e9xg5hvdtq9ktye4rtmt",
				"delegationFee": 31250
			},
			{
				"nodeID": "NodeID-QB2RK6w4RQeAqNvxpGmCAxEDkYHE1uZut",
				"rewardAddress": "X-dijets1g7cxf7g8ml90zg399ewgpg8rw6c9r380xf7xsl",
				"delegationFee": 20000
			},
			{
				"nodeID": "NodeID-8mQWMPdp44nPAxsjvbyfgTDhUpmWUVN3u",
				"rewardAddress": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"delegationFee": 20000
			},
			{
				"nodeID": "NodeID-LbqzrSMKAssm5Ds2bJ8LPwymnGJbrybCN",
				"rewardAddress": "X-dijets19qeq9fg568ztpace8mdecq4jd9rx5dj64fc4sf",
				"delegationFee": 20000
			},
			{
				"nodeID": "NodeID-Nh98zaRumrBz3K5i3SQG1T3WrhJv1Jxmb",
				"rewardAddress": "X-dijets15j0a4x07vz8fahf4racp6mg5h7pw4l8kw9pa7y",
				"delegationFee": 20000
			},
			{
				"nodeID": "NodeID-4tsvZR6LEaVc8mbc5K4Kusji21YZDdk39",
				"rewardAddress": "X-dijets174rcm4tva8f7t4m7u2nxe7a236nv9pcwhdmt5v",
				"delegationFee": 20000
			},
			{
				"nodeID": "NodeID-KnSQaB5tDy1rgZWHFdCHSwTTdNvRNkLBP",
				"rewardAddress": "X-dijets1g7cxf7g8ml90zg399ewgpg8rw6c9r380xf7xsl",
				"delegationFee": 20000
			},
			{
				"nodeID": "NodeID-BFa1padLXBj7VHa2JYvYGzcTBPQGjPhUy",
				"rewardAddress": "X-dijets1g7cxf7g8ml90zg399ewgpg8rw6c9r380xf7xsl",
				"delegationFee": 20000
			},
			{
				"nodeID": "NodeID-4B4rc5vdD1758JSBYL1xyvE5NHGzz6xzH",
				"rewardAddress": "X-dijets1g7cxf7g8ml90zg399ewgpg8rw6c9r380xf7xsl",
				"delegationFee": 20000
			},
			{
				"nodeID": "NodeID-EDESh4DfZFC15i613pMtWniQ9arbBZRnL",
				"rewardAddress": "X-dijets1g7cxf7g8ml90zg399ewgpg8rw6c9r380xf7xsl",
				"delegationFee": 20000
			},
			{
				"nodeID": "NodeID-CZmZ9xpCzkWqjAyS7L4htzh5Lg6kf1k18",
				"rewardAddress": "X-dijets1g7cxf7g8ml90zg399ewgpg8rw6c9r380xf7xsl",
				"delegationFee": 20000
			},
			{
				"nodeID": "NodeID-CTtkcXvVdhpNp6f97LEUXPwsRD3A2ZHqP",
				"rewardAddress": "X-dijets1g7cxf7g8ml90zg399ewgpg8rw6c9r380xf7xsl",
				"delegationFee": 20000
			},
			{
				"nodeID": "NodeID-84KbQHSDnojroCVY7vQ7u9Tx7pUonPaS",
				"rewardAddress": "X-dijets1g7cxf7g8ml90zg399ewgpg8rw6c9r380xf7xsl",
				"delegationFee": 20000
			},
			{
				"nodeID": "NodeID-JjvzhxnLHLUQ5HjVRkvG827ivbLXPwA9u",
				"rewardAddress": "X-dijets1g7cxf7g8ml90zg399ewgpg8rw6c9r380xf7xsl",
				"delegationFee": 20000
			},
			{
				"nodeID": "NodeID-4CWTbdvgXHY1CLXqQNAp22nJDo5nAmts6",
				"rewardAddress": "X-dijets1g7cxf7g8ml90zg399ewgpg8rw6c9r380xf7xsl",
				"delegationFee": 20000
			}
		],
		"cChainGenesis": "{\"config\":{\"chainId\":43113,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"0100000000000000000000000000000000000000\":{\"code\":\"0x7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80631e010439146042578063b6510bb314606e575b600080fd5b605c60048036036020811015605657600080fd5b503560b1565b60408051918252519081900360200190f35b818015607957600080fd5b5060af60048036036080811015608e57600080fd5b506001600160a01b03813516906020810135906040810135906060013560b6565b005b30cd90565b836001600160a01b031681836108fc8690811502906040516000604051808303818888878c8acf9550505050505015801560f4573d6000803e3d6000fd5b505050505056fea26469706673582212201eebce970fe3f5cb96bf8ac6ba5f5c133fc2908ae3dcd51082cfee8f583429d064736f6c634300060a0033\",\"balance\":\"0x0\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "hi mom"
	}`

	// FujiParams are the params used for the fuji testnet
	FujiParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                 units.MilliDjtx,
			CreateAssetTxFee:      10 * units.MilliDjtx,
			CreateSubnetTxFee:     100 * units.MilliDjtx,
			CreateBlockchainTxFee: 100 * units.MilliDjtx,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement: .8, // 80%
			MinValidatorStake: 1 * units.Djtx,
			MaxValidatorStake: 3 * units.MegaDjtx,
			MinDelegatorStake: 1 * units.Djtx,
			MinDelegationFee:  20000, // 2%
			MinStakeDuration:  24 * time.Hour,
			MaxStakeDuration:  365 * 24 * time.Hour,
			RewardConfig: reward.Config{
				MaxConsumptionRate: .12 * reward.PercentDenominator,
				MinConsumptionRate: .10 * reward.PercentDenominator,
				MintingPeriod:      365 * 24 * time.Hour,
				SupplyCap:          100 * units.MegaDjtx,
			},
		},
	}
)
