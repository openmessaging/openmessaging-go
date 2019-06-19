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

package consumer

import (
	. "github.com/openmessaging-go/openmessaging"
)

type Context interface {
	/**
	 * Acknowledges the specified and consumed message, which is related to this {@code MessageContext}.
	 * <p>
	 * Messages that have been received but not acknowledged may be redelivered.
	 *
	 * @return error if the consumer fails to acknowledge the messages due to some internal error.
	 */
	Ack() error
}

type MessageListener interface {
	/**
	 * Callback method to receive incoming messages.
	 * <p>
	 * A message listener should handle different types of {@code Message}.
	 *
	 * @param message the received message object.
	 * @param context the context delivered to the consume thread.
	 */
	OnReceived(message Message, context Context) error
}
