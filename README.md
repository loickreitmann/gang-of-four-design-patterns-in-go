# Design Patterns in Go

**Design Patterns** are a set of 23 classic design patterns in software development, described in the book *Design Patterns: Elements of Reusable Object-Oriented Software*. The book was authored by Erich Gamma, Richard Helm, Ralph Johnson, and John Vlissides, who are collectively known as the **Gang of Four**. They described these patterns as solutions to common problems in software design. They are divided into three categories: Creational, Structural, and Behavioral.

The following design patterns are foundational in object-oriented software development and are widely used to create flexible, reusable, and maintainable code.

## A. Creational Patterns

#### 1. [Abstract Factory](creational/abstract-factory/main.go)
Provides an interface for creating families of related or dependent objects without specifying their concrete classes.

#### 2. [Builder](creational/builder/main.go)
Separates the construction of a complex object from its representation so that the same construction process can create different representations.

#### 3. [Factory Method](creational/factory-method/main.go)
Defines an interface for creating an object but lets subclasses alter the type of objects that will be created.

#### 4. [Prototype](creational/prototype/main.go)
Specifies the kinds of objects to create using a prototypical instance and creates new objects by copying this prototype.

#### 5. [Singleton](creational/singleton/main.go)
Ensures a class has only one instance and provides a global point of access to it.

## B. Structural Patterns

#### 1. [Adapter](structural/adapter/main.go)
Converts the interface of a class into another interface that clients expect, allowing classes to work together that couldn't otherwise because of incompatible interfaces.

#### 2. [Bridge](structural/bridge/main.go)
Separates an object's interface from its implementation so the two can vary independently.

#### 3. [Composite](structural/composite/main.go)
Composes objects into tree structures to represent part-whole hierarchies, allowing clients to treat individual objects and compositions uniformly.

#### 4. [Decorator](structural/decorator/main.go)
Attaches additional responsibilities to an object dynamically, providing a flexible alternative to subclassing for extending functionality.

#### 5. [Facade](structural/facade/main.go)
Provides a unified interface to a set of interfaces in a subsystem, making the subsystem easier to use.

#### 6. [Flyweight](structural/flyweight/main.go)
Uses sharing to support large numbers of fine-grained objects efficiently.

#### 7. [Proxy](structural/proxy/main.go)
Provides a surrogate or placeholder for another object to control access to it.

## C. Behavioral Patterns

#### 1. [Chain of Responsibility](behavioral/chain-of-responsibility/main.go)
Passes a request along a chain of handlers, allowing multiple objects a chance to handle the request.

#### 2. [Command](behavioral/command/main.go)
Encapsulates a request as an object, thereby allowing for parameterization of clients with queues, requests, and operations.

#### 3. [Interpreter](behavioral/interpreter/main.go)
Defines a representation for a language's grammar and uses an interpreter to interpret sentences in the language.

#### 4. [Iterator](behavioral/iterator/main.go)
Provides a way to access the elements of an aggregate object sequentially without exposing its underlying representation.

#### 5. [Mediator](behavioral/mediator/main.go)
Defines an object that encapsulates how a set of objects interact, promoting loose coupling.

#### 6. [Memento](behavioral/memento/main.go)
Captures and externalizes an object's internal state without violating encapsulation, so the object can be restored to this state later.

#### 7. [Observer](behavioral/observer/main.go)
Defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.

#### 8. [State](behavioral/state/main.go)
Allows an object to alter its behavior when its internal state changes, appearing as if the object changed its class.

#### 9. [Strategy](behavioral/strategy/main.go)
Defines a family of algorithms, encapsulates each one, and makes them interchangeable. Strategy lets the algorithm vary independently from clients that use it.

#### 10. [Template Method](behavioral/template-method/main.go)
Defines the skeleton of an algorithm in an operation, deferring some steps to subclasses.

#### 11. [Visitor](behavioral/visitor/main.go)
Represents an operation to be performed on elements of an object structure, allowing new operations to be defined without changing the classes of the elements on which it operates.
