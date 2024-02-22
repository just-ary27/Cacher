# Cacher
A simple library in Go that does a basic implementation of the few most popular caching algorithms.

The caching eviction policies implemented are the following:
- FIFO (First in first out)
- LIFO (Last in first out)
- LRU (Least Recently Used)

You can make your own policy by embedding the `EvictionPolicy` base struct in your policy.

## Examples
Please note that in each example, the respective cache variable has been typed as `interfaces.Cache`to prevent access to the internal state in the program.

- FIFO
    - [Code](https://github.com/just-ary27/Cacher/blob/main/examples/fifo_example.go)
    - Output: <br/>
![image](https://github.com/just-ary27/Cacher/assets/76696648/920e1957-1596-4686-b6e4-c9013e13c6fc)
![image](https://github.com/just-ary27/Cacher/assets/76696648/b1d0627a-efea-49d3-b467-8cbe363f27de)
![image](https://github.com/just-ary27/Cacher/assets/76696648/df3f95cf-ebed-4ba9-9882-cff380889764)

- LIFO
    - [Code](https://github.com/just-ary27/Cacher/blob/main/examples/lifo_example.go)
    - Output: <br/>
  ![image](https://github.com/just-ary27/Cacher/assets/76696648/1db920ac-645f-4239-835f-232a32b03bcc)


- LRU
    - [Code](https://github.com/just-ary27/Cacher/blob/main/examples/lru_example.go)
    - Output: <br/>
  ![image](https://github.com/just-ary27/Cacher/assets/76696648/1e5a6163-d176-4016-a48e-c5ac1c4925af)


## Tech Stack
- go v1.21

<div align=center>
    <hr>
    <p>Made with ❤️ by 
        <a href="https://justary27.web.app">
            <img src="https://user-images.githubusercontent.com/76696648/176264414-6b9a9549-cb25-41f3-9e00-d857bd8bd7cf.svg" align=center height=40>
        </a>
    </p>
</div>
