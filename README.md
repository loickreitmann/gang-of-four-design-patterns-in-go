# Design Patters in Go

**Design Patterns** are a set of 23 classic design patterns in software development, described in the book *Design Patterns: Elements of Reusable Object-Oriented Software*. The book was authored by Erich Gamma, Richard Helm, Ralph Johnson, and John Vlissides, who are collectively known as the **Gang of Four**. They described these patterns as solutions to common problems in software design. They are divided into three categories: Creational, Structural, and Behavioral.

## Creational Patterns

#### 1. [Abstract Factory](creational/abstract-factory/main.go)
Provides an interface for creating families of related or dependent objects without specifying their concrete classes.

#### 2. [Builder](creational/builder/main.go)
Separates the construction of a complex object from its representation so that the same construction process can create different representations.

#### 3. Factory Method
Defines an interface for creating an object but lets subclasses alter the type of objects that will be created.

#### 4. Prototype
Specifies the kinds of objects to create using a prototypical instance and creates new objects by copying this prototype.

#### 5. Singleton
Ensures a class has only one instance and provides a global point of access to it.