---
title: 主語・目的語入門
icon: ./latin.ico
---

# [こばのページ](../index.html)/[ラテン語](index.html)/主語・目的語入門

(Status: ![WIP](https://progress-bar.dev/60?title=WIP))

- 対格について書く: 80%
- 主格について書く: 40%

# このページについて
ラテン語の主語と目的語について、簡単に入門します。ラテン語文を読んだ時に大体目的語っぽいものを判別できるようになることが目標です。

ラテン語作文ができるようになることは目標としていないので、その話は一切しません。

# 主格・対格について
ラテン語には 5 個の格 (主格・属格・与格・対格・奪格) があり[^where-is-vocative]、文章の単語はすべてこれらのうちのいずれかに分類できます。これらのうちのいずれになるかによって、単語の語尾がわずかに変化することになります。

これらのうち出現頻度が高く文章の基本となるのは主格・対格の 2 種類です。この文書ではこれらにフォーカスし、ラテン語文の核が大まかに分かるようになることを目標とします。

[^where-is-vocative]: これら 5 個の格に加え呼格も含めて 6 個とみなすこともあります。しかし呼格はほとんどの場合主格と同じですので、ここでは 5 個ということにします。

- 主格 (nominative): 主語の格。基本的に「〜は」と訳す。辞書にもこの形で載っている。
- 対格 (accusative): 目的語の格。基本的に「〜を」と訳す。一部の前置詞 (ad, in (…へ), propter, per, secundum など) の後の単語もこれになる。

# 大まかな判別方法

## 対格、単数
単数の場合、-m で終わっていたら多くの場合で対格です。

- 例 (単数):
  - -am で終わるパターン: [**vi**am](nom-acc-intro/noun-via), [**glō**riam](nom-acc-intro/noun-gloria), [**dex**teram](nom-acc-intro/noun-dextera), [**vī**tam](nom-acc-intro/noun-vita)
  - -um で終わるパターン: **De**um, [pec**cā**tum](nom-acc-intro/noun-peccatum), [**Do**minum](nom-acc-intro/noun-Dominus), [**Spī**ritum](nom-acc-intro/noun-spiritus)
  - -em で終わるパターン: [**pā**cem](nom-acc-intro/noun-pax), [dēprecāti**ō**nem](nom-acc-intro/noun-deprecatio), [**Pa**trem](nom-acc-intro/noun-pater), [fac**tō**rem](nom-acc-intro/noun-factor), [**re**quiem](nom-acc-intro/noun-requies)
  - -im, -om で終わるパターン: ほとんどありません。見かけたら天然記念物として大切にしましょう。

- 例文:
  - **Dō**nā **nō**bīs [**pa**cem](nom-acc-intro/noun-pax), et **dō**nā **e**īs [**re**quiem](nom-acc-intro/noun-requies).
    - 訳: `私たちに平和を与えろ、そして彼らに安息を与えろ。`
    - pacem が「平和を」、requiem が「安息を」という意味です。これらは対格です。
  - **Grā**tiās **a**gimus **ti**bi **prop**ter **mag**nam [**glō**riam](nom-acc-intro/noun-gloria) **tu**am. (ミサ通常文 Gloria より)
    - 訳: `あなたの大いなる栄光を理由として、私たちはあなたに感謝する。`
    - propter は「〜ゆえに」を意味する前置詞で、対格を取ります。magnam は「大きい」を意味する形容詞 magnus の女性・単数・対格、tuam は「あなたの」を意味する形容詞 tuus の女性・単数・対格です。形容詞で修飾された語が対格になる場合、形容詞も含め全てが対格になります。(そのためここでの magnam, gloriam, tuam すべてが対格です。)

注意 1: -um で終わっても対格ではない場合があります。
複数属格 (「…たちの」) というパターンです。
-ārum, -ōrum のパターンは割と発見しやすいです (in saecula saeculōrum など)。それ以外の -um で終わるパターン (vīsibilium omnium (すべての見えるものたちの) など) は個別対応しかないでしょう。
- 例: Et exspectō [resurrectiōnem](nom-acc-intro/noun-resurrectio) [mortuōrum](nom-acc-intro/noun-mortuus). (ミサ通常文 Credo より)
    - 訳: `そして私は、死者たちの復活を待ち望む。`
    - resurrectiōnem は「復活を」を意味し (単数・対格)、mortuōrum は「死者たちの」を意味します (複数・属格)。


注意 2: -m で終わらないのに対格である場合があります。
中性名詞と呼ばれるカテゴリーで、これらの名詞は主格と対格が<u>**常に同じ形**</u>です。このルールは先ほどの「単数対格は -m で終わる」という法則よりはるかに強いです。
- 例: Dā [**rō**bur](nom-acc-intro/noun-robur).
  - 訳: `力を与えろ。`
  - rōbur は中性名詞なので、単数対格は rōburem などにはならず rōbur のままです。


## 対格、複数
複数の場合、単数の場合より区別が難しいです。-ās, -ōs, -ēs あたりの語尾で区別しましょう。幸いなことに頻度もそんなに多くないです。

- 例 (複数):
  - -ās で終わるパターン: Scrip**tū**rās, Pro**phē**tās, [**grā**tiās](nom-acc-intro/noun-gratia)
  - -ōs で終わるパターン: [**vī**vōs](nom-acc-intro/noun-vivus), [**mor**tuōs](nom-acc-intro/noun-mortuus)
  - -ēs で終わるパターン: **ho**minēs
  - -a で終わるパターン (中性名詞): [**bo**na](nom-acc-intro/noun-bonum), [**ma**la](nom-acc-intro/noun-malum)
  - -ūs で終わるパターン: [**ma**nūs](nom-acc-intro/noun-manus)
    - 頻度はそれほど多くないはずです。筆者は in **ma**nūs **tu**ās com**men**dō **spī**ritum **me**um (`私はあなたの両手へと私の魂を委ねる`) でしか見たことがありません。

- 例文:
  - [**Grā**tiās](nom-acc-intro/noun-gratia) **a**gimus **ti**bi **prop**ter **mag**nam **glō**riam **tu**am. (ミサ通常文 Gloria より)
    - 訳: `あなたの大いなる栄光を理由として、私たちはあなたに感謝する。`
    - grātiās は gratia (感謝) の複数・対格です。agimus は「私たちはなす」なので、直訳は「私たちは感謝たちをなす」です。
  - jūdi**cā**re [**vī**vōs](nom-acc-intro/noun-vivus) et [**mor**tuōs](nom-acc-intro/noun-mortuus) (ミサ通常文 Credo より)
    - 訳: `生者たちと死者たちを裁くために`
  - [**Ma**la](nom-acc-intro/noun-malum) **nos**tra **pel**le, [**bo**na](nom-acc-intro/noun-bonum) **cunc**ta **pos**ce. ([歌詞](https://koba-e964.gitbook.io/choral-music-latin-trans/avemarisstella/3))
    - 訳: `私たちの悪を打ち破れ、すべての善を願え。`
    - mala は [malum (悪)](nom-acc-intro/noun-malum) の複数・対格で、bona は [bonum (善)](nom-acc-intro/noun-bonum) の複数・対格です。どちらも中性名詞なので主格も対格も同じ形ですが、ここでは文脈で目的語であろう (つまり対格であろう) と推測できます。

注意: -ās で終わるパターンには抽象名詞の -itās で終わるパターンがあります。(英語の -ity に相当。) これらは全然複数対格ではないので注意しましょう。
- 例: cāritās (charity, 隣人愛), vēritās (verity, 真実), vānitās (vanity, 空虚さ), castitās (chastity, 純潔), gravitās (gravity, 深刻さ)

## 主格
色々な形をとります。辞書に載っている形なので、見かけた時に辞書で調べることは比較的簡単です。

以下に、主格の単数形と複数形を載せます。英語でもこのような複数形になる単語があるので、馴染みのある方もいらっしゃるでしょう。

||-us|-a|-um|その他|
|--|--|--|--|--|
|単数|-us|-a|-um|その他|
|複数|-ī|-ae|-a|-ēs|

例です。

||-us|-a|-um|その他|
|--|--|--|--|--|
|単数|focus, alumnus|formula, alumna|stadium|index, vertex|
|複数|focī, alumnī|formulae, alumnae|stadia|indicēs, verticēs|


## 主格、単数
-us, -a, -um は言うに及ばず、さらに -s とか -x で終わっていたら主格です。ただし -is で終わっていたら主格でないことの方が多いです。

- 例 (単数):
  - -us で終わるパターン: [**Do**minus](nom-acc-intro/noun-Dominus), [**Spī**ritus](nom-acc-intro/noun-spiritus)
  - -um で終わるパターン: [pec**cā**tum](nom-acc-intro/noun-peccatum), mystērium, sacrāmentum
    - 「中性名詞は主格と対格が同じ形」という話を思い出しましょう。
  - -s で終わるパターン: [**re**quies](nom-acc-intro/noun-requies)
  - -x で終わるパターン: [pax](nom-acc-intro/noun-pax)
  - -a で終わるパターン: [**vi**a](nom-acc-intro/noun-via), [**glō**ria](nom-acc-intro/noun-gloria), [**dex**tera](nom-acc-intro/noun-dextera), [**vī**ta](nom-acc-intro/noun-vita)

- 例文:
  - [**Do**minus](nom-acc-intro/noun-Dominus) **tē**cum.
    - 訳: `主はあなたと共にある。`
    - [**Do**minus](nom-acc-intro/noun-Dominus) は「主」を意味します。
    - **tē**cum は「あなたと共に」を意味する単語です。
  - Be**ā**tī **quō**rum [**vi**a](nom-acc-intro/noun-via) **in**tegra est.
    - 訳: `道が汚れていない者たちは幸福である。`
    - [**vi**a](nom-acc-intro/noun-via) は単数・主格です。「道が」を意味します。integra は形容詞 integer (完全な、汚れていない) の女性・単数・主格であって、via にかかります。
  - In **ter**rā [pax](nom-acc-intro/noun-pax) ho**mi**nibus **bo**nae volun**tā**tis. (ミサ通常文 Gloria より)
    - 訳: `地において、平和が善意の人間たちにありますように。`
    - [pax](nom-acc-intro/noun-pax) は単数・主格です。「平和が」を意味し、この文の主語です。
  - Tū **sō**lus [**Do**minus](nom-acc-intro/noun-Dominus) (es). (ミサ通常文 Gloria より)
    - 訳: `あなただけが主である。`
    - [**Do**minus](nom-acc-intro/noun-Dominus) は「主」を意味します。
    - **sō**lus は「唯一の」を意味する形容詞です。
    - es は「…である」を意味する単語です。ここでは tū (あなた) が主語で **Do**minus は補語ですが、英語などと違い es の補語は主格にします。英語では It's me. や It's him. などのように、be 動詞の補語は目的格 (ラテン語の対格に相当) にするのでした。
  - **Sanc**ta Ma**rī**a, **mā**ter **De**ī.
    - 訳: `聖なるマリーアよ、神の母よ。`
    - 主格は呼びかけにも使います。呼びかけに使われる格を「呼格」と呼ぶ流儀もありますが、多くの場合で主格と同じです。
  - [**Do**mine](nom-acc-intro/noun-Dominus), **Fī**lī ūni**ge**nite.
    - 訳: `主よ、唯一生み出された子よ。`
    - 呼びかけの時だけ主格から外れることもあります。**Do**mine は [**Do**minus](nom-acc-intro/noun-Dominus) (主) の呼びかけの形、**Fī**lī は **Fī**lius (子) の呼びかけの形、ūni**ge**nite は ūni**ge**nitus (唯一生み出された、お気に入りの) の呼びかけの形です。こういった場合については、呼びかけであることは明らかなことが多いので、形が多少変わることもあると覚えておけば良いはずです。
