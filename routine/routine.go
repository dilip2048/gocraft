/*
 * // Copyright (c) 2018-2020 by ConverseNow Technologies Inc. All rights reserved.
 * // DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
 * //
 * // All information contained herein is proprietary and confidential to ConverseNow
 * // Technologies Inc. Any use, reproduction, or disclosure without the written
 * // permission of ConverseNow Technologies Inc. is prohibited.
 * //
 * // 	Author : dilip.yadav@conversenow.ai
 *
 */

package routine

import (
"fmt"
"sync"
)

func routine() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go printHelloWorld(&wg, i)
	}
	wg.Wait()
}

func printHelloWorld(wg *sync.WaitGroup, flag int) {
	defer wg.Done()
	fmt.Printf("%v: Hello World\n", flag)
}
