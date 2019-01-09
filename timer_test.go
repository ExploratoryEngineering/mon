package mon

//
//Copyright 2018 Telenor Digital AS
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
import "testing"

func TestTimer(t *testing.T) {
	time1 := newHistogramCounter("time.1")
	time2 := newHistogramCounter("time.2")
	tk := NewTimer()

	tk.Begin(time1)
	tk.End()

	tk.Begin(nil)

	tk.Begin(time1)
	tk.Begin(nil)

	tk.Begin(time2)
	tk.Begin(time2)
	tk.End()

	tk.End()
}
