# 设计模式

## 创建型模式

- 简单工厂模式(Simple Factory)

    go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。 NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。
    
- 工厂方法模式(Factory Method)

    工厂方法模式的实质是"定义一个创建对象的接口，但让实现这个接口的类来决定实例化哪个类。工厂方法让类的实例化推迟到子类中进行。"

- 抽象工厂模式(Abstract Factory)

    抽象工厂模式用于生成产品族的工厂，所生成的对象是有关联的。如果抽象工厂退化成生成的对象无关联则成为工厂函数模式。

- 创建者模式(Builder)

    一个对象的创建十分复杂，为了区分构建过程和使用过程，因此分开。使用一个Director类进行对象的创建，Builder规定了这个创建过程。

- 原型模式(ProtoType)

    用于创建重复的对象，同时又能保证性能。当直接创建对象的代价比较大时，则采用这种模式。

- 单例模式(Singleton)

    使用懒惰模式的单例模式，使用双重检查加锁保证线程安全

## 结构型模式

- 外观模式(Facade)

- 适配器模式(Adapter)

- 代理模式(Proxy)

- 组合模式(Composite)

- 享元模式(FlyWeight)

- 装饰模式(Decorator)

- 桥模式(Bridge)

## 行为型模式

- 中介者模式(Mediator)

- 观察者模式(Observer)

- 命令模式(Command)

- 迭代器模式(Iterator)

- 模板方法模式(Template Method)

- 策略模式(Strategy)

- 状态模式(State)

- 备忘录模式(Memento)

- 解释器模式(Interpreter)

- 职责链模式(Chain of Responsibility)

- 访问者模式(Visitor)