---
layout: post
title:  "有限状态机"
date:   2021-01-14 12:00:00 +0800
tags:   todo
---

* category
{:toc}



[golang looplab/fsm](https://github.com/looplab/fsm)

[Javascript Finite State Machine](https://github.com/jakesgordon/javascript-state-machine)


# States and Transitions 状态与转换

![matter state machine](https://raw.githubusercontent.com/jakesgordon/javascript-state-machine/master/examples/matter.png)

A state machine consists of a set of **states**, e.g:

状态机一般会包含一组 **状态** ，比如：

  * 固态 solid
  * 液态 liquid
  * 气态 gas

.. and a set of **transitions**, e.g:

.. 和一组 **转换方法（动作）**, e.g:

  * 融化 melt
  * 冰冻 freeze
  * 蒸发 vaporize
  * 凝结 condense


```javascript
  var fsm = new StateMachine({
    init: 'solid',
    transitions: [
      { name: 'melt',     from: 'solid',  to: 'liquid' },
      { name: 'freeze',   from: 'liquid', to: 'solid'  },
      { name: 'vaporize', from: 'liquid', to: 'gas'    },
      { name: 'condense', from: 'gas',    to: 'liquid' }
    ]
  });

  fsm.state;             // 'solid'
  fsm.melt();
  fsm.state;             // 'liquid'
  fsm.vaporize();
  fsm.state;             // 'gas'
```

## Multiple states for a transition 一个动作与多种不同状态间的转换

![wizard state machine](https://raw.githubusercontent.com/jakesgordon/javascript-state-machine/master/examples/wizard.png)

If a transition is allowed `from` multiple states then declare the transitions with the same name:

可以给相同的动作名称定义不同的状态转换逻辑：

```javascript
  { name: 'step',  from: 'A', to: 'B' }, //当前状态为 A 时，执行 step 动作后，状态转换为 B ;
  { name: 'step',  from: 'B', to: 'C' }, //当前状态为 B 时，执行 step 动作后，状态转换为 C ;
  { name: 'step',  from: 'C', to: 'D' }  //当前状态为 C 时，执行 step 动作后，状态转换为 D ;
```

If a transition with multiple `from` states always transitions `to` the same state, e.g:

同一个动作名称，也可以从不同的状态转换到固定的一个状态：

```javascript
  { name: 'reset', from: 'B', to: 'A' },
  { name: 'reset', from: 'C', to: 'A' },
  { name: 'reset', from: 'D', to: 'A' }
```

... then it can be abbreviated using an array of `from` states:

上述定义也能像下面这样更简便的方法定义：

```javascript
  { name: 'reset', from: [ 'B', 'C', 'D' ], to: 'A' }
```

Combining these into a single example:

将以上逻辑组合到一个示例代码中就是这样：

```javascript
  var fsm = new StateMachine({
    init: 'A',
    transitions: [
      { name: 'step',  from: 'A',               to: 'B' },
      { name: 'step',  from: 'B',               to: 'C' },
      { name: 'step',  from: 'C',               to: 'D' },
      { name: 'reset', from: [ 'B', 'C', 'D' ], to: 'A' }
    ]
  })
```

This example will create an object with 2 transition methods:

以上示例能创建一个拥有两个成员方法（动作）的（状态机）对象：

  * `fsm.step()`
  * `fsm.reset()`

The `reset` transition will always end up in the `A` state, while the `step` transition
will end up in a state that is dependent on the current state.

   `reset` 方法总是把状态改成 A ；
而 `step` 方法则会根据上一个状态进行转换。


## Wildcard Transitions 使用通配符执行状态转换

If a transition is appropriate from **any** state, then a wildcard `*` `from` state can be used:

如果一个动作可在 **任意** 状态时执行， 可以使用 `*` 通配符表示；

```javascript
  var fsm = new StateMachine({
    transitions: [
      // ...
      { name: 'reset', from: '*', to: 'A' }
    ]
  });
```

## Conditional Transitions 根据条件执行状态转换

A transition can choose the target state at run-time by providing a function as the `to` attribute:

给 `to` 属性赋值为函数，就能在运行时根据当前状态，动态调整要转换的目标状态：

```javascript
  var fsm = new StateMachine({
    init: 'A',
    transitions: [
      { name: 'step', from: '*', to: function(n) { return increaseCharacter(this.state, n || 1) } }
    ]
  });

  fsm.state;      // A
  fsm.step();
  fsm.state;      // B
  fsm.step(5);
  fsm.state;      // G

  // helper method to perform (c = c + n) on the 1st character in str
  function increaseCharacter(str, n) {
    return String.fromCharCode(str.charCodeAt(0) + n);
  }
```

The `allStates` method will only include conditional states once they have been seen at run-time:

`allStates` 能返回所有在运行时出现过的状态：

```javascript
  fsm.state;        // A
  fsm.allStates();  // [ 'A' ]
  fsm.step();
  fsm.state;        // B
  fsm.allStates();  // [ 'A', 'B' ]
  fsm.step(5);
  fsm.state;        // G
  fsm.allStates();  // [ 'A', 'B', 'G' ]
```

## GOTO - Changing State Without a Transition 不经过状态转换，直接改变状态

You can use a conditional transition, combined with a wildcard `from`, to implement
arbitrary `goto` behavior:



```javascript
  var fsm = new StateMachine({
    init: 'A'
    transitions: [
      { name: 'step', from: 'A', to: 'B'                      },
      { name: 'step', from: 'B', to: 'C'                      },
      { name: 'step', from: 'C', to: 'D'                      },
      { name: 'goto', from: '*', to: function(s) { return s } }
    ]
  })

  fsm.state;     // 'A'
  fsm.goto('D');
  fsm.state;     // 'D'
```

A full set of [Lifecycle Events](https://raw.githubusercontent.com/jakesgordon/javascript-state-machine/master/docs/lifecycle-events.md) still apply when using `goto`.

