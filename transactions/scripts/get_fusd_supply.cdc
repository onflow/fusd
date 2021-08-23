// This script returns the total supply of FUSD.

import FUSD from 0xFUSDADDRESS

pub fun main(): UFix64 {
    return FUSD.totalSupply
}
