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

type ServiceLifeState int

const (
	/**
	 * Service has been initialized.
	 */
	INITIALIZED ServiceLifeState = iota

	/**
	 * Service in starting.
	 */
	STARTING

	/**
	 * Service in running.
	 */
	STARTED

	/**
	 * Service is stopping.
	 */
	STOPPING

	/**
	 * Service has been stopped.
	 */
	STOPPED
)

func (this ServiceLifeState) String() string {
	switch this {
	case INITIALIZED:
		return "INITIALIZED"
	case STARTING:
		return "STARTING"
	case STARTED:
		return "STARTED"
	case STOPPING:
		return "STOPPING"
	case STOPPED:
		return "STOPPED"
	default:
		return "UNKNOWN"
	}
}

type ServiceLifecycle interface {
	/**
	 * Used for startup or initialization of a service endpoint. A service endpoint instance will be in a ready state
	 * after this method has been completed.
	 */
	Start() error

	/**
	 * Notify a service instance of the end of its life cycle. Once this method completes, the service endpoint could be
	 * destroyed and eligible for garbage collection.
	 */
	Stop() error

	/**
	 * Used for get service current state, for execution of some operations is dependent on the current service state.
	 *
	 * @return This service current state {@link ServiceLifeState}
	 */
	CurrentState() (ServiceLifeState, error)
}
