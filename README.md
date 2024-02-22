# Cacher
A simple library in Go that does a basic implementation of the few most popular caching algorithms.

The caching eviction policies implemented are the following:
- FIFO (First in first out)
- LIFO (Last in first out)
- LRU (Least Recently Used)

You can make your own policy by embedding the `EvictionPolicy` base struct in your policy.

## Examples
Please note that cache has been typed as `interfaces.Cache`to prevent access to internal state in the program.

- FIFO
    - Code: 
    - Output:

- LIFO
    - Code: 
    - Output:

- LRU
    - Code:
    - Output:

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
