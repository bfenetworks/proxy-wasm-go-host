/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1

import (
	"github.com/bfenetworks/proxy-wasm-go-host/proxywasm/common"
	v1 "github.com/bfenetworks/proxy-wasm-go-host/proxywasm/v1"
)

func copyIntoInstance(instance common.WasmInstance, value string, retPtr int32, retSize int32) v1.WasmResult {
	addr, err := instance.Malloc(int32(len(value)))
	if err != nil {
		return v1.WasmResultInvalidMemoryAccess
	}

	err = instance.PutMemory(addr, uint64(len(value)), []byte(value))
	if err != nil {
		return v1.WasmResultInvalidMemoryAccess
	}

	err = instance.PutUint32(uint64(retPtr), uint32(addr))
	if err != nil {
		return v1.WasmResultInvalidMemoryAccess
	}

	err = instance.PutUint32(uint64(retSize), uint32(len(value)))
	if err != nil {
		return v1.WasmResultInvalidMemoryAccess
	}

	return v1.WasmResultOk
}

func copyBytesIntoInstance(instance common.WasmInstance, value []byte, retPtr int32, retSize int32) v1.WasmResult {
	addr, err := instance.Malloc(int32(len(value)))
	if err != nil {
		return v1.WasmResultInvalidMemoryAccess
	}

	err = instance.PutMemory(addr, uint64(len(value)), value)
	if err != nil {
		return v1.WasmResultInvalidMemoryAccess
	}

	err = instance.PutUint32(uint64(retPtr), uint32(addr))
	if err != nil {
		return v1.WasmResultInvalidMemoryAccess
	}

	err = instance.PutUint32(uint64(retSize), uint32(len(value)))
	if err != nil {
		return v1.WasmResultInvalidMemoryAccess
	}

	return v1.WasmResultOk
}

func getContextHandler(instance common.WasmInstance) v1.ContextHandler {
	if v := instance.GetData(); v != nil {
		if im, ok := v.(v1.ContextHandler); ok {
			return im
		}
	}

	return nil
}

func getImportHandler(instance common.WasmInstance) v1.ImportsHandler {
	if ctx := getContextHandler(instance); ctx != nil {
		if im := ctx.GetImports(); im != nil {
			return im
		}
	}

	return &v1.DefaultImportsHandler{}
}
