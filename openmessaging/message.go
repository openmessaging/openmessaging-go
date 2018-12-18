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

type Message interface {
	/**
		 * Returns all the system header fields of the {@code Message} object as a {@code KeyValue}.
		 *
		 * @return the system headers of a {@code Message}
		 */
	Headers() (Headers, error)

	/**
	 * Returns all the customized user header fields of the {@code Message} object as a {@code KeyValue}.
	 *
	 * @return the user properties of a {@code Message}
	 */
	Properties() (KeyValue, error)

	/**
	 * Get data from message body
	 *
	 * @return message body
	 * @throws OMSMessageFormatException if the message body cannot be assigned to the specified type
	 */
	GetData() ([]byte, error)

	/**
	 * Set data to message body
	 *
	 * @param data set message body in binary stream
	 */
	SetData(data []byte) error
}
