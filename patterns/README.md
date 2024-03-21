# Design patterns
are typical solutions to common problems in software design. They represent best practices used by
experienced object-oriented software developers.

### Design patterns are categorized into three main groups:

1. **Creational Patterns**: These patterns provide various object creation mechanisms, which increase flexibility and
   reuse of existing code.
    - **Singleton**: Ensures that a class has only one instance and provides a global point of access to it.
    - **Factory Method**: Defines an interface for creating an object, but lets subclasses alter the type of objects
      that will be created.
    - **Abstract Factory**: Provides an interface for creating families of related or dependent objects without
      specifying their concrete classes.
    - **Builder**: Allows for the construction of complex objects step by step. It’s often used when an object needs to
      be initialized with many possible configurations.
    - **Prototype**: Creates new objects by copying an existing object, known as the prototype.

2. **Structural Patterns**: These patterns explain how to assemble objects and classes into larger structures while
   keeping these structures flexible and efficient.
    - **Adapter (or Wrapper)**: Allows objects with incompatible interfaces to collaborate by wrapping its own interface
      around that of an already existing class.
    - **Composite**: Composes objects into tree structures to represent part-whole hierarchies. It lets clients treat
      individual objects and compositions of objects uniformly.
    - **Proxy**: Provides a placeholder for another object to control access to it. This is used for lazy loading,
      controlling the access, or logging, for example.
    - **Flyweight**: Reduces the cost of creating and manipulating a large number of similar objects.
    - **Bridge**: Separates an object’s abstraction from its implementation so that the two can vary independently.

3. **Behavioral Patterns**: These patterns are all about class's objects communication. They help in defining how
   objects interact in a way that increases flexibility in carrying out communication.
    - **Observer**: Allows one object to notify other objects of any state changes, typically used for implementing
      distributed event handling systems.
    - **Strategy**: Defines a family of algorithms, encapsulates each one, and makes them interchangeable. Strategy lets
      the algorithm vary independently from clients that use it.
    - **Command**: Turns a request into a stand-alone object that contains all information about the request. This
      transformation lets you parameterize methods with different requests, delay or queue a request’s execution, and
      support undoable operations.
    - **State**: Allows an object to alter its behavior when its internal state changes. The object will appear to
      change its class.
    - **Template Method**: Defines the skeleton of an algorithm in the superclass but lets subclasses override specific
      steps of the algorithm without changing its structure.

Each group and pattern serves a specific purpose, offering a structured approach to solving design problems in software
development.