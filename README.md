pubsub benchmark between Go and Node.js
=========


##requirements
Go: redis client(fixed from [hoisie's redis.go](https://github.com/alphazero/Go-Redis))

Node.js: 

    `npm install -g hredis`

##system info

Linux krazy 3.4.4-2-ARCH #1 SMP PREEMPT Sun Jun 24 18:59:47 CEST 2012 x86_64 GNU/Linux

processor   : 3

vendor_id   : GenuineIntel

cpu family  : 6

model       : 37

model name  : Intel(R) Core(TM) i5 CPU       M 520  @ 2.40GHz

stepping    : 2

microcode   : 0x9

cpu MHz     : 2400.000

cache size  : 3072 KB



##benchmark statistics
###100 clients
####Go:
2285.816579 task-clock                  #    0.023 CPUs utilized          

27,459 context-switches                 #    0.012 M/sec                  

4,265 CPU-migrations                    #    0.002 M/sec                  

142,385 page-faults                     #    0.062 M/sec                  

3,800,092,301 cycles                    #    1.662 GHz                     [83.31%]

2,525,479,877 stalled-cycles-frontend   #   66.46% frontend cycles idle    [82.59%]

1,425,696,558 stalled-cycles-backend    #   37.52% backend  cycles idle    [68.15%]

2,426,433,455 instructions              #    0.64  insns per cycle        

1.04  stalled cycles per insn [89.79%]

567,501,478 branches                    #  248.271 M/sec                   [85.34%]

20,909,990 branch-misses                #    3.68% of all branches         [83.77%]

100.235526142 seconds time elapsed

####Node:
9474.336083 task-clock                  #    0.094 CPUs utilized          

54,295 context-switches                 #    0.006 M/sec                  

8,984 CPU-migrations                    #    0.948 K/sec                  

457,762 page-faults                     #    0.048 M/sec                  

23,682,864,275 cycles                   #    2.500 GHz                     [83.15%]

14,983,946,055 stalled-cycles-frontend  #   63.27% frontend cycles idle    [83.69%]

9,961,776,331 stalled-cycles-backend    #   42.06% backend  cycles idle    [67.12%]

18,546,817,673 instructions             #    0.78  insns per cycle        

0.81  stalled cycles per insn [84.45%]

3,495,682,550 branches                  #  368.963 M/sec                   [82.92%]

128,459,342 branch-misses               #    3.67% of all branches         [83.39%]

100.268050890 seconds time elapsed

##300 clients
####Go:
10484.851531 task-clock                 #    0.035 CPUs utilized          

75,854 context-switches                 #    0.007 M/sec                  

11,308 CPU-migrations                   #    0.001 M/sec                  

873,176 page-faults                     #    0.083 M/sec                  

17,816,076,446 cycles                   #    1.699 GHz                     [83.65%]

11,556,723,029 stalled-cycles-frontend  #   64.87% frontend cycles idle    [83.56%]

6,622,424,927 stalled-cycles-backend    #   37.17% backend  cycles idle    [67.76%]

12,970,109,944 instructions             #    0.73  insns per cycle        

0.89  stalled cycles per insn [86.08%]

2,805,685,026 branches                  #  267.594 M/sec                   [82.28%]

85,749,038 branch-misses                #    3.06% of all branches         [83.46%]

301.485277275 seconds time elapsed

####Node:
65203.089838 task-clock                 #    0.216 CPUs utilized          

2,376,256 context-switches              #    0.036 M/sec                  

134,276 CPU-migrations                  #    0.002 M/sec                  

1,540,857 page-faults                   #    0.024 M/sec                  

160,028,622,698 cycles                  #    2.454 GHz                     [83.16%]

116,724,130,577 stalled-cycles-frontend #   72.94% frontend cycles idle    [83.72%]

76,554,624,978 stalled-cycles-backend   #   47.84% backend  cycles idle    [66.83%]

93,263,532,979 instructions             #    0.58  insns per cycle        

1.25  stalled cycles per insn [83.74%]

17,980,457,917 branches                 #  275.761 M/sec                   [83.02%]

686,829,923 branch-misses               #    3.82% of all branches         [83.38%]

301.789098647 seconds time elapsed

##500 clients
####Go:

14657.905211 task-clock                 #    0.031 CPUs utilized          

276,221 context-switches                #    0.019 M/sec                  

12,559 CPU-migrations                   #    0.857 K/sec                  

1,128,474 page-faults                   #    0.077 M/sec                  

26,168,154,602 cycles                   #    1.785 GHz                     [82.37%]

16,793,404,475 stalled-cycles-frontend  #   64.17% frontend cycles idle    [82.73%]

9,470,198,837 stalled-cycles-backend    #   36.19% backend  cycles idle    [67.67%]

19,090,185,661 instructions             #    0.73  insns per cycle        

0.88  stalled cycles per insn [86.50%]

4,110,371,712 branches                  #  280.420 M/sec                   [84.46%]

132,114,611 branch-misses               #    3.21% of all branches         [83.71%]

474.435301811 seconds time elapsed


####Node:

175533.859312 task-clock                #    0.299 CPUs utilized          

6,887,182 context-switches              #    0.039 M/sec                  

329,808 CPU-migrations                  #    0.002 M/sec                  

3,045,929 page-faults                   #    0.017 M/sec                  

442,918,334,628 cycles                  #    2.523 GHz                     [83.36%]

337,066,430,244 stalled-cycles-frontend #   76.10% frontend cycles idle    [83.34%]

229,580,670,733 stalled-cycles-backend  #   51.83% backend  cycles idle    [66.90%]

226,626,576,488 instructions            #    0.51  insns per cycle        

1.49  stalled cycles per insn [83.62%]

43,966,825,492 branches                 #  250.475 M/sec                   [83.25%]

1,652,911,838 branch-misses             #    3.76% of all branches         [83.31%]

502.714263363 seconds time elapsed

##注意事项
1. redis提示too many open connections,尝试ulimit -m 4096或更大
2. 测试go和node.js实现前，重启一下redis-server


##结论
Go性能在用户量增大的情况下，性能会比node.js下的性能有明显的提高
