/* Licensed to the Apache Software Foundation (ASF) under one or more
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

package producer

import (
	. "github.com/openmessaging-go/openmessaging"
	. "github.com/openmessaging-go/openmessaging/interceptor"
)

type Producer interface {
	ServiceLifecycle
	MessageFactory
	/**
	 * Sends a message to the specified destination synchronously, the destination should be preset to {@link
	 * Message#headers()}, other header fields as well.
	 *
	 * @param message a message will be sent.
	 * @return the successful {@code SendResult}.
	 * @return error when have no authority to send messages to a given destination.
	 * @return error when an invalid message is specified.
	 * @return error when the given timeout elapses before the send operation completes.
	 * @return error when have no given destination in the server.
	 * @return error when the {@code Producer} fails to send the message due to some internal error.
	 */
	Send(message Message) (SendResult, error)

	/**
	 * Sends a message to the specified destination asynchronously, the destination should be preset to {@link
	 * Message#headers()}, other header fields as well.
	 * <p>
	 * The returned {@code *chan SendResult} will have the result once the operation completes, and the SendResult {@code
	 * SendResult} will be set to the channel, either because the operation was successful or because of an error.
	 *
	 * @param message a message will be sent.
	 * @return the {@code *chan SendResult} of an asynchronous message send operation.
	 * @see Future
	 * @see FutureListener
	 */
	SendAsync(message Message) (*chan SendResult, error)

	/**
	 * <p>
	 * There is no {@code SendResult} related or error returned. The calling thread doesn't care about the
	 * send result and also have no context to get the result.
	 *
	 * @param message a message will be sent.
	 */
	SendOneway(message Message)

	/**
	 * <p>
	 * Send batch messages to server.
	 *
	 * @param messages messages to be sent.
	 */
	SendBatch(messages []Message) error

	/**
	 * Adds a {@code ProducerInterceptor} to intercept send operations of producer.
	 *
	 * @param interceptor a producer interceptor.
	 */
	AddInterceptor(interceptor ProducerInterceptor) error

	/**
	 * Removes a {@code ProducerInterceptor}.
	 *
	 * @param interceptor a producer interceptor will be removed.
	 */
	RemoveInterceptor(interceptor ProducerInterceptor) error

	/**
	 * Sends a transactional message to the specified destination synchronously, the destination should be preset to
	 * {@link Message#headers()}, other header fields as well.
	 * <p>
	 * A transactional send result will be exposed to consumer if this prepare message send success, and then, you can
	 * execute your local transaction, when local transaction execute success, users can use {@link
	 * TransactionalResult#commit()} to commit prepare message,otherwise can use {@link TransactionalResult#rollback()}
	 * to roll back this prepare message.
	 * </p>
	 *
	 * @param message a prepare transactional message will be sent.
	 * @return the successful {@code SendResult}.
	 * @return error when have no authority to send messages to a given destination.
	 * @return error when an invalid message is specified.
	 * @return error when the given timeout elapses before the send operation completes.
	 * @return error when have no given destination in the server.
	 * @return error when the {@code Producer} fails to send the message due to some internal error.
	 * @return error when used normal producer which haven't register {@link TransactionStateCheckListener}.
	 */
	Prepare(message Message) (TransactionalResult, error)
}
