__Project 0__

[Office hours video](https://drive.google.com/file/d/1N5EEeBIGG_LbOPOeplu6_7maHthOt-F9/view)

* Objective is to implement an associative array (ADT). An associative array offers the following interface
	- `allocate(size)`
	- `put(key, value)`
	- `get(key, values_array)`
	- `erase(key)`
	- `deallocate()`

* Actual implementation is in the form of a hashtable.

* The hashtable has a separate chaining structure, with fat nodes to take advantage of sequential access.

* The hashtable has to be re-sized if it becomes too large or if the chains become too long.

* Steps:
	- Implement the API
	- Fatness of nodes
	- Resizing of hashtable

