---
title: 主語・目的語入門
icon: ./latin.ico
---

# [こばのページ](../index.html)/[ラテン語](index.html)/主語・目的語入門

WIP: ![WIP](https://progress-bar.dev/40?title=WIP)

- 対格について書く: 80%
- 主格について書く: 0%

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
  - -um で終わるパターン: **De**um, pec**cā**tum, **Do**minum, **Spī**ritum
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
中性名詞と呼ばれるカテゴリーで、これらの名詞は主格と対格が常に同じ形です。このルールは先ほどの「単数対格は -m で終わる」という法則よりはるかに強いです。
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
