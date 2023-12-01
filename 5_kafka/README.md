# Kafka

#### Kafka 在 consumer-side 實現擴展性的方式：

將 consumer 組織成 consumer groups，
同一個的 consumer group 內的多個 consumer 之間並行處理 topic 和 partition。
可以透過加入更多的 consumer 來水平擴展。


#### 當 consumer 過多，數據過少時會導致 Kafka consumer 效能變慢：

如果將過多的 consumer 分配給相同的主題，而主題內的分區太少，Kafka 可能會將同一分區劃分給多個 consumer。
這可能會減慢速度，因為它增加了 Kafka cluster 必須執行的計算開銷，決定將每個主題的哪些部分交付給每個 consumer。

#### 解決方案：

在決定將多少 consumer 加入 consumer組時，避免加入多於分區數的 consumer。


