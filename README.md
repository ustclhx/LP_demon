# LP_demon

## Introduction

LP_demon is a golang libary for causal inference, aiming to provides a series of unified interface to realize various stages of inference, including **making assumption in Causal Graphical Models**, **indentifying by do-calculus**, and **estimating the effect given control variables**.

The theoretical basis mainly come Judea Pearl, the Turing award winner in 2011, and the estimation stage partly based on the potential outcomes framework of Rubin.

The name of LP_demon stands for the assumption that causality in the macro world should be deterministic. And if we have all the right assumptions (based on our knowledge) and enough data(off course it's impossible), we might be able to make a Laplace demon to analyze and predict the macro world.(It can not predict the free will of human, but we can tell it what we well do)

For now, it just a very early stage of development for LP-demon, it's just a minimum executavle prototype.

The version 0.1.0 is released on 2020.1.16.

## Requirements

You need to install [Go1.13 or later first](https://golang.org/doc/install).

And there are 3 dependent library of  v0.1.0(as it's a earliest beta version, the dependencies may be changed in the future):

* [gophersat](https://github.com/crillab/gophersat)  
* [gonum](https://github.com/gonum/gonum)  
* [goml](https://github.com/cdipaolo/goml)

## Example 

[An easiest example](https://github.com/ustclhx/LP_demon/blob/master/example.go) of v0.1.0

## Wish List

1.short-term:     
* Improve the algorithm for identify(backdoor and frontdoor)    
* More estimation method
* Frontdoor-based estimate   
* Mulitvariate treatment / Non-binary treatment    
  
2.long-term:
* refute
* pdag
* discover
    
3.engineering:  
* python api
* distributed
