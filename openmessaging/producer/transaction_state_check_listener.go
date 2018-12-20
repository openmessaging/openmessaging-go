/* Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License") you may not use this file except in compliance with
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

package producer

import (
	. "github.com/openmessaging-go/openmessaging"
)

type TransactionalContext interface {
	/**
	 * Commits a transaction.
	 */
	Commit() error

	/**
	 * Rolls back a transaction.
	 */
	Rollback() error

	/**
	 * Unknown transaction status, may be this transaction still on going.
	 */
	Unknown() error
}

type TransactionStateCheckListener interface {
	/**
	* Checks the status of the local transaction branch.
	*
	* @param message the associated message.
	* @param context the check context.
	*/
	Check(message Message, context TransactionalContext) error
}
