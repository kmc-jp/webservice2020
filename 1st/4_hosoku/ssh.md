# 常態化 SSH
<!-- TOC -->

- [常態化 SSH](#常態化-ssh)
    - [概要](#概要)
    - [準備](#準備)
    - [おわり](#おわり)

<!-- /TOC -->
## 概要
皆さんに作ってもらうプログラムは常態化させるために最終的にサーバ上で動かす必要があります。そこでどのようにサーバにあげるのかを簡単に説明します。

## 準備
KMCのサーバにログインし、Golangをインストールします。

まずはログイン。
```sh
$ ssh KMC_USERNAME@kmc.gr.jp 
PassWord: KMCのUNIXアカウントパスワード
Linux ringo(省略)
USER_NAME@ringo:~$ 
```

Goのインストール
```sh
$ wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
$ tar -C ~/ -xzf go1.14.4.linux-amd64.tar.gz 

# YOUR_USER_NAMEを埋める必要があるので注意！
$ echo PATH=$PATH:/home/YOUR_USER_NAME/go/bin >> .bashrc
$ echo GOPATH=/home/YOUR_USER_NAME/go >> .bashrc

```

これでインストールが完了しました。実行するためにローカルにあるスクリプトをサーバにコピーします。ログアウトするまえにプログラムを格納するディレクトリを作成してきましょう。

```sh
$ mkdir Program
$ Ctrl+D(control+D)
ログアウト
```

それでは自分のプログラムを送信します。

リモートへ/からのコピーは一般に次の文法で記述されます。

```sh
$ scp 送信元 送信先
```
但し、リモートのディレクトリの前にはその先の`ユーザ名@アドレス`を入力しましょう。今回は次のようになります。

```
$ scp -r 自分のプログラムの入ったディレクトリ KMC_USERNAME@kmc.gr.jp:~/Program/
Password:YOUR_KMC_PASSWORD
```
これでコピーできました。あとはコンパイルして実行するだけです。

```sh
$ ssh KMC_USERNAME@kmc.gr.jp
Password:YOUR_KMC_PASSWORD
$ cd Program/先程のプログラムのディレクトリ名/
$ go build
$ nohup ./先程のプログラムのディレクトリ名 &
```
これで実行されました。

止めるときは次のようにします。

```sh
$ ps -x
  PID TTY      STAT   TIME COMMAND
(省略)
数値  ？       なにか hh:mm ./先程のプログラムのディレクトリ名
(省略)
$ kill 数値
```

以上です。
## おわり
毎度パスワード打つの面倒に感じたあなたはKMCのハンドブック、解説記事/UNIX/sshの項目を参照しましょう。あなたの助けになるはずです。

それでは本題の
[Git／GitHubに入りましょう！](../git/readme.md)