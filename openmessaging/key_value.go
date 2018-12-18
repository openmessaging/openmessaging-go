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

type KeyValue interface {
	/**
     * Inserts or replaces {@code short} value for the specified key.
     *
     * @param key the key to be placed into this {@code KeyValue} object
     * @param value the value corresponding to <tt>key</tt>
     */
	PutInt16(key string, value int16) (KeyValue, error)

	/**
	 * Inserts or replaces {@code int} value for the specified key.
	 *
	 * @param key the key to be placed into this {@code KeyValue} object
	 * @param value the value corresponding to <tt>key</tt>
	 */
	PutInt(key string, value int) (KeyValue, error)

	/**
	 * Inserts or replaces {@code long} value for the specified key.
	 *
	 * @param key the key to be placed into this {@code KeyValue} object
	 * @param value the value corresponding to <tt>key</tt>
	 */
	PutInt64(key string, value int64) (KeyValue, error)

	/**
	 * Inserts or replaces {@code double} value for the specified key.
	 *
	 * @param key the key to be placed into this {@code KeyValue} object
	 * @param value the value corresponding to <tt>key</tt>
	 */
	PutFloat64(key string, value float64) (KeyValue, error)

	/**
	 * Inserts or replaces {@code String} value for the specified key.
	 *
	 * @param key the key to be placed into this {@code KeyValue} object
	 * @param value the value corresponding to <tt>key</tt>
	 */
	Put(key string, value string) (KeyValue, error)

	/**
	 * Searches for the {@code short} property with the specified key in this {@code KeyValue} object. If the key is not
	 * found in this property list, zero is returned.
	 *
	 * @param key the property key
	 * @return the value in this {@code KeyValue} object with the specified key value
	 * @see #put(String, short)
	 */
	GetShort(key string) (int16, error)

	/**
	 * Searches for the {@code int} property with the specified key in this {@code KeyValue} object. If the key is not
	 * found in this property list, zero is returned.
	 *
	 * @param key the property key
	 * @return the value in this {@code KeyValue} object with the specified key value
	 * @see #put(String, int)
	 */
	GetInt(key string) (int16, error)

	/**
	 * Searches for the {@code long} property with the specified key in this {@code KeyValue} object. If the key is not
	 * found in this property list, zero is returned.
	 *
	 * @param key the property key
	 * @return the value in this {@code KeyValue} object with the specified key value
	 * @see #put(String, long)
	 */
	GetLong(key string) (int64, error)

	/**
	 * Searches for the {@code double} property with the specified key in this {@code KeyValue} object. If the key is
	 * not found in this property list, zero is returned.
	 *
	 * @param key the property key
	 * @return the value in this {@code KeyValue} object with the specified key value
	 * @see #put(String, double)
	 */
	GetDouble(key string) (float64, error)

	/**
	 * Searches for the {@code String} property with the specified key in this {@code KeyValue} object. If the key is
	 * not found in this property list, {@code null} is returned.
	 *
	 * @param key the property key
	 * @return the value in this {@code KeyValue} object with the specified key value
	 * @see #put(String, String)
	 */
	GetString(key string) (string, error)

	/**
	 * Returns a {@link Set} view of the keys contained in this {@code KeyValue} object.
	 * <p>
	 * The set is backed by the {@code KeyValue}, so changes to the set are reflected in the @code KeyValue}, and
	 * vice-versa.
	 *
	 * @return the key set view of this {@code KeyValue} object.
	 */
	 //todo
	//Set<String> KeySet()

	/**
	 * Tests if the specified {@code String} is a key in this {@code KeyValue}.
	 *
	 * @param key possible key
	 * @return <code>true</code> if and only if the specified key is in this {@code KeyValue}, <code>false</code>
	 * otherwise.
	 */
	ContainsKey(key string) (bool, error)
}
