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
package openmessaging

import (
	. "github.com/openmessaging-go/openmessaging/consumer"
	. "github.com/openmessaging-go/openmessaging/manager"
	. "github.com/openmessaging-go/openmessaging/producer"
)

type MessagingAccessPoint interface {
	/**
	 * Returns the target OMS specification version of the specified vendor implementation.
	 *
	 * @return the OMS version of implementation
	 * @see OMS#specVersion
	 */
	Version() (string, error)

	/**
	 * Returns the attributes of this {@code MessagingAccessPoint} instance.
	 * <p>
	 * There are some standard attributes defined by OMS for {@code MessagingAccessPoint}:
	 * <ul>
	 * <li> {@link OMSBuiltinKeys#ACCESS_POINTS}, the specified access points.
	 * <li> {@link OMSBuiltinKeys#DRIVER_IMPL}, the fully qualified class name of the specified MessagingAccessPoint's
	 * implementation, the default value is {@literal io.openmessaging.<driver_type>.MessagingAccessPointImpl}.
	 * <li> {@link OMSBuiltinKeys#REGION}, the region the resources reside in.
	 * <li> {@link OMSBuiltinKeys#ACCOUNT_ID}, the ID of the specific account system that owns the resource.
	 * </ul>
	 *
	 * @return the attributes
	 */
	Attributes() (KeyValue, error)

	/**
	 * Creates a new {@code Producer} for the specified {@code MessagingAccessPoint}.
	 *
	 * @return the created {@code Producer}
	 * @return error if the {@code MessagingAccessPoint} fails to handle this request due to some internal
	 * error
	 * @return error if have no authority to create a producer.
	 */
	CreateProducer() (Producer, error)

	/**
	 * Creates a new transactional {@code Producer} for the specified {@code MessagingAccessPoint}, the producer is able
	 * to respond to requests from the server to check the status of the transaction.
	 *
	 * @param transactionStateCheckListener transactional check listener {@link TransactionStateCheckListener}
	 * @return the created {@code Producer}
	 * @return error if the {@code MessagingAccessPoint} fails to handle this request due to some internal
	 * error
	 * @return error if have no authority to create a producer.
	 */
	CreateTransactionProducer(transactionStateCheckListener TransactionStateCheckListener) (Producer, error)

	/**
	 * Creates a new {@code PushConsumer} for the specified {@code MessagingAccessPoint}. The returned {@code Consumer}
	 * isn't bind to any queue, uses {@link Consumer#bindQueue(String, MessageListener)} to bind queues.
	 *
	 * @return the created {@code PushConsumer}
	 * @return error if the {@code MessagingAccessPoint} fails to handle this request due to some internal
	 * error
	 * @return error if have no authority to create a consumer.
	 */
	CreateConsumer() (Consumer, error)

	/**
	 * Gets a lightweight {@code ResourceManager} instance from the specified {@code MessagingAccessPoint}.
	 *
	 * @return the resource manger
	 * @return error if the {@code MessagingAccessPoint} fails to handle this request due to some internal
	 * error
	 * @return error if have no authority to obtain a resource manager.
	 */
	ResourceManager() (ResourceManager, error)
}
