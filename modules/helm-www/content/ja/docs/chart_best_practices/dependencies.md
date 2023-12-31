---
title: "依存関係"
description: "チャートの依存関係についての解説です。"
weight: 4
---

このセクションでは、`Chart.yaml`に宣言された`依存関係`のベストプラクティスについて
解説します。

## バージョン

可能であれば、正確なバージョンを固定するよりバージョンの範囲を使用してください。デフォルト
では、パッチレベルのバージョンマッチを使用することが推奨されています:

```yaml
version: ~1.2.3
```

これはバージョン`1.2.3`とそのバッチバージョンにマッチします。つまり、
`~1.2.3`は`>= 1.2.3, < 1.3.0`と同等になります。

完全なバージョンマッチのシンタックスについては、[semver
documentation](https://github.com/Masterminds/semver#checking-version-constraints)
を参照してください。

### プレリリースバージョン

上記のバージョンの制約は、プレリリースバージョンにはマッチしません。
例えば、`version: ~1.2.3`は`version: ~1.2.4`にマッチしますが、`version: ~1.2.3-1`
にはマッチしません。
以下はパッチレベルのマッチングだけでなく、プレリリースにも対応します:

```yaml
version: ~1.2.3-0
```

### リポジトリURL

可能であれば、`https://`のリポジトリURLを使用し、その後、`http://`のURLを使用してください。

リポジトリが、そのインデックスファイルに追加されている場合、リポジトリの名前がURLの別名
として使われます。`alias:`もしくは`@`の後に、リポジトリ名を付けて利用してください。

ファイルのURL（`file://...`）は固定されたデプロイパイプラインによって組み立てられた
チャートのための"特殊なケース"と考えられます。

[downloader plugins]({{< ref path="/docs/topics/plugins.md#downloader-plugins" lang="en" >}})を使用する場合、
URLのスキームはプラグインに固有になります。注意としては、チャートの利用者はその依存関係を
ビルドもしくは更新するために、スキームをサポートするプラグインをインストールする必要があります。

`repository`フィールドが空白の場合、Helmは依存関係の管理操作を行うことができません。
その場合、Helmはその依存関係が`name`で指定した名前と同じ名前の、
`charts`フォルダーのサブディレクトリにあると推測します。

## コンディションとタグ

_任意の_ 依存関係には、コンディションもしくはタグを追加するべきです。

コンディションの好ましい書き方:

```yaml
condition: somechart.enabled
```

ここでは`somechart`が依存関係のあるチャートの名前です。

複数のサブチャート（依存関係）が共に任意もしくは取り替え可能な機能
を提供する場合、これらのチャートは同じタグを共有するべきです。

例えば、`nginx`と`memcached`の両方がチャート内で、メインアプリケーション用に
パフォーマンスの最適化を提供している場合、そして、その機能を使用する際、その両方が
必要になる場合、このようなタグのセクションを含めるべきでしょう:

```yaml
tags:
  - webaccelerator
```

これにより、利用者はその機能のオンとオフを一つのタグで切り替えることができます。
