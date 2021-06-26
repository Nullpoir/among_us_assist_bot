# Among us進行補助bot概要
これはdiscord内で動くAmong usの進行を補助するアプリケーションです。
vc接続時に問題となる一斉mute切り替えについて比較的環境を選ばない形で解決をしています。

# 使用方法

## discord側の設定
1. ゲームをする際に使用するdiscordのカテゴリーを一つ作成します。
1. vcチャンネルとしてgameとmute、テキストチャンネルとしてbot操作もしくはチャットという名前でテキストチャンネルを作成します。なぜbot操作もしくはチャットで操作ができるのかというとdiscord内にてvcを使えないメンバーがテキストチャンネルで議論をするというユースケースを想定しているためです。その心配がないのであればbot操作という名前にするのが好ましいと思います。
1. 以上で設定は完了です。

## botの起動
ある程度のPCスキルを有する前提で、解説してまいります。
現状このbotを一般利用できる形(botを招待することで、利用可能になる形態)での提供は考えておりません。理由は処理速度として問題ない形で提供するのにコストがかかるためです。そのためお手数ですが、自力でbotを起動するようにしてください。環境差が大きいので概要だけまとめます。
1. dockerを利用するかローカル環境にredisをインストールする。dockerを利用する場合はdokcer compose up -dをご利用いただけるようにymlを書いてあります。
1. BOTをdiscordの公式サイトで登録して、アクセストークンなどを取得してください。
1. settigsディレクトリの.exmple_envを参考に同ディレクトリに.envを作成してください。
1. `yarn install`で必要なパッケージをインストールし、`yarn build`からの`yarn start`でサーバーが起動します。

以上になります。

個別の環境設定についてissueを立てて質問頂いても構いませんが、回答ができるかはわかりませんので予めご了承ください。

# bot操作方法
bot操作もしくはチャットにて文字を送信することで操作できます。

1. m

mと一文字bot操作もしくはチャットで送信すると、送信したテキストチャンネルが所属しているカテゴリーに存在するgameという名前のVCにいるメンバーのミュート/アンミュートを切り替えられます。

2. c

ミュート/アンミュートすべきかどうかの情報をクリアします。このbotはミュート/アンミュートすべきかどうかを独自に判断することが出来ます。しかしアルゴリズムは単純なので、実際の状態とずれるかもしれません。そういうときにリセットをかけるコマンドがcです。

## 今後のメンテナンス
開発者個人のdiscordサーバーの要望に応じて機能追加をいたします。git pull origin masterで最新コードにアップデートされますが、変更によっては動かなくなる可能性があります。予めご了承ください。

