## Tests cli and build

```bash
$ ./build.sh
$ ./bin/sample c1 -n netWork -r 2
beforeAction is called()..
command1 action is running.. network : network and repeat : 2
0 -> network
1 -> network
afterAction is called()..
$ ./bin/sample command2 --network NetWork -r 3
beforeAction is called()..
command2 action is running.. network : NETWORK and repeat : 3
0 -> NETWORK
1 -> NETWORK
2 -> NETWORK
afterAction is called()..
```  