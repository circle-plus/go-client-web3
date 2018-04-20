/********************************************************************************
   This file is part of go-web3.
   go-web3 is free software: you can redistribute it and/or modify
   it under the terms of the GNU Lesser General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-web3 is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Lesser General Public License for more details.
   You should have received a copy of the GNU Lesser General Public License
   along with go-web3.  If not, see <http://www.gnu.org/licenses/>.
*********************************************************************************/

/**
 * @file block.go
 * @authors:
 *   Jérôme Laurens <jeromelaurens@gmail.com>
 * @date 2017
 */

package dto

import (
	// "encoding/json"
	// "fmt"
	// "strconv"

	"github.com/regcostajr/go-web3/complex/types"
)

type Block struct {
	Number     types.ComplexIntResponse `json:"number"`
	Hash       string                   `json:"hash"`
	ParentHash string                   `json:"parentHash"`
	Author     string                   `json:"author,omitempty"`
	Miner      string                   `json:"miner,omitempty"`
	Size       types.ComplexIntResponse `json:"size"`
	GasUsed    types.ComplexIntResponse `json:"gasUsed"`
	Nonce      types.ComplexIntResponse `json:"nonce"`
	Timestamp  types.ComplexIntResponse `json:"timestamp"`
}
