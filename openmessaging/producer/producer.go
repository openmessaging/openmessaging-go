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
     * @throws OMSSecurityException when have no authority to send messages to a given destination.
     * @throws OMSMessageFormatException when an invalid message is specified.
     * @throws OMSTimeOutException when the given timeout elapses before the send operation completes.
     * @throws OMSDestinationException when have no given destination in the server.
     * @throws OMSRuntimeException when the {@code Producer} fails to send the message due to some internal error.
     */
	Send(message Message) (SendResult, error)

	/**
	 * Sends a message to the specified destination asynchronously, the destination should be preset to {@link
	 * Message#headers()}, other header fields as well.
	 * <p>
	 * The returned {@code Promise} will have the result once the operation completes, and the registered {@code
	 * FutureListener} will be notified, either because the operation was successful or because of an error.
	 *
	 * @param message a message will be sent.
	 * @return the {@code Promise} of an asynchronous message send operation.
	 * @see Future
	 * @see FutureListener
	 */
	 //todo
	SendAsync(message Message) error

	/**
	 * <p>
	 * There is no {@code Promise} related or {@code RuntimeException} thrown. The calling thread doesn't care about the
	 * send result and also have no context to get the result.
	 *
	 * @param message a message will be sent.
	 */
	SendOneway(message Message) error

	/**
	 * <p>
	 * Send batch messages to server.
	 *
	 * @param messages messages to be sent.
	 */
	 //todo
	//SendBatch(List<Message> messages) error

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
	 * @throws OMSSecurityException when have no authority to send messages to a given destination.
	 * @throws OMSMessageFormatException when an invalid message is specified.
	 * @throws OMSTimeOutException when the given timeout elapses before the send operation completes.
	 * @throws OMSDestinationException when have no given destination in the server.
	 * @throws OMSRuntimeException when the {@code Producer} fails to send the message due to some internal error.
	 * @throws OMSTransactionException when used normal producer which haven't register {@link TransactionStateCheckListener}.
	 */
	Prepare(message Message) (TransactionalResult, error)
}
