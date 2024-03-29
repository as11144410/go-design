# go-design
go设计模式
- 五大基本原则
    - 单一功能原则
    - 开闭原则
    - 里氏替换原则
    - 接口隔离原则
    - 依赖反转原则
- 设计模式分类
    - 创建型（单例模式、简单工厂模式、工厂方法模式、抽象工厂模式、建造者模式、原型模式）
    - 结构型（代理模式、适配器模式、装饰模式、桥接模式、组合模式、享元模式、外观模式）
    - 行为型（观察者模式、模版方法模式、命令模式、状态模式、职责链模式、解释器模式、中介者模式、访问者模式、策略模式、备忘录模式、迭代器模式）
- 简单工厂
    - 优点：实现了解藕
    - 缺点：违背开闭原则
    - 适合：创建的对象比较少
- 工厂方法模式
    - 优点：保持了简单工厂模式的优点，克服了简单工厂模式的缺点
    - 缺点：在添加新产品时，在一定程度上增加了系统复杂度
    - 适合：客户端不需要知道具体产品类的类名，只需要知道所对应的工厂即可
- 装饰器模式
    - 一种动态的往一个类中添加新的行为的设计模式
    - 就功能而言，装饰模式相比生成子类更为灵活，这样可以给某个对象而不是整个类添加一些功能
    - 它是一种对象结构型模式
    - 优点：
        - 可以通过一种动态的方式来扩展一个对象的功能
        - 可以使用多个具体的装饰类来装饰同一对象，增加其功能
        - 具体组件类与具体装饰类可以独立变化，符合开闭原则
    - 缺点：
        - 对于多次装饰的对象，易出错，排除困难
        - 产品很多具体装饰类，增加系统复杂度及理解成本
    - 适合场景
        - 需要给一个对象增加功能，这些功能可以动态的撤销
        - 需要给一批兄弟类增加或者改装功能

- 策略模式
    - 策略模式作为一种软件设计模式，指对象有某个行为，但是在不同的场景中，该行为有不同的实现算法
    - 比如每个人都要交个人所得税，但是在美国和中国交，就有不同的算税方法
    - 策略模式是一种对象行为模式
    - 优点：
        - 对开闭原则的完美支持
        - 避免使用多重条件转移语句
        - 提供了管理相关的算法族的办法
    - 缺点：
        - 客户端必须知道所有策略类，并自己决定使用哪一个策略类
        - 策略模式将造成产生很多策略类
    - 适合：
        - 一个系统需要动态的在几种算法或者行为中选择一种
        - 多个类区别仅在于他们的行为或算法不同的场景