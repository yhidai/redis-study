## 確認手順

### レプリケーション構成のRedis、Sentinelを作成する

### マスター、レプリカ用のラベルを付与する
レプリケーション構成のRedis、SentinelのPodが全てActiveになった後、ラベルを付与する。
下記の例ではredis-replication-2がマスターだったため、マスター用のラベルをつけている。

```sh
kubectl label pods redis-replication-2 redis-role=master
kubectl label pods redis-replication-0 redis-role=replica
kubectl label pods redis-replication-1 redis-role=replica
kubectl label pods redis-replication-3 redis-role=replica
kubectl get pods -l role=replication --show-labels
```

### NodePortのサービスを作成する


### クライアント環境からアクセスする



