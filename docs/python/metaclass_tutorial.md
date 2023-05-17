---
title: Tutorial Metaclass
tags:
    - python
    - tutorial
---

Original link [Python Metaclass Tutorial (with Examples)](https://coderslegacy.com/python-metaclass-tutorial/)

*In this Python tutorial, we will discuss the concept of a “Metaclass”, and how it can be used effectively.*

Metaclasses are a powerful feature in Python that allow you to create classes dynamically at runtime. In Python, everything is an object, including classes. When you define a class in Python, you are actually creating an object of type type. This means that you can use metaclasses to create custom classes with their own behaviors and attributes.

In a nutshell, a metaclass is a class that creates other classes. When you define a class with a metaclass, you are essentially telling Python to use that metaclass to create the class object. Metaclasses can be used to add new functionality to classes, modify the behavior of existing classes, and perform validation on class attributes.

The importance of metaclasses in Python programming cannot be overstated. They are a powerful tool that can be used to create more dynamic and flexible programs. With metaclasses, you can customize the way classes are created and add new features to your classes that would be difficult or impossible to achieve with normal class inheritance.

This is just one example of the many ways that metaclasses can be used to customize class creation in Python. In the following sections of this Python tutorial, we’ll explore the basics of a metaclass, as well as more advanced features and use cases.

## The Basics of Metaclasses in Python

In order to understand metaclasses in Python, it’s important to first understand the ```type()``` function. The ```type()``` function is the built-in metaclass in Python, and is used to create new classes.

When you define a new class in Python, you are actually calling the ```type()``` function with three arguments: the name of the class, the tuple of base classes, and a dictionary containing the attributes of the class. For example:

```python
class MyClass:
    pass
```

Is equivalent to:

```python
MyClass = type('MyClass', (), {})
```

In this example, we are calling the type() function with the arguments 'MyClass', an empty tuple for the base classes, and an empty dictionary for the attributes.

Now that we understand the type() function, let’s take a look at how we can create custom metaclasses. To create a custom metaclass, we simply define a new class that inherits from type. For example:

```python
class MyMeta(type):
    pass
```

In this example, we are defining a new class called ```MyMeta``` that inherits from type. This means that ```MyMeta``` is itself a metaclass, and can be used to create new classes.

When creating a custom metaclass, we can define two special methods: ```__new__()``` and ```__init__()```.

The ```__new__()``` method is called when the metaclass is used to create a new class object, and is responsible for creating and returning the new class object. The ```__init__()``` method is called after the new class object has been created, and is responsible for initializing the class attributes.

Here’s an example of a custom metaclass that adds a new class attribute:

```python
class MyMeta(type):
    def __new__(cls, name, bases, attrs):
        attrs['custom_attribute'] = 'Hello, world!'
        return super().__new__(cls, name, bases, attrs)
 
class MyClass(metaclass=MyMeta):
    pass

print(MyClass.custom_attribute)  # Output: "Hello, world!"
```

In this example, we define a new metaclass called MyMeta that adds the custom_attribute to every new class it creates. Then, we define a new class called MyClass that uses MyMeta as its metaclass. When we create an instance of MyClass, we can see that it automatically has the custom_attribute.

In the next section, we will explore these two functions (new and init) in greater detail with some new examples.

## Creating Custom Metaclasses in Python

Now that we have an understanding of the basics of metaclasses and how to define them, let’s explore some more examples of how to use metaclasses and their associated methods.

### new() Function

The new() method is responsible for creating and returning a new class object. It takes four parameters:

* cls: the metaclass being called
* name: the name of the class being created
* bases: a tuple of base classes
* attrs: a dictionary of class attributes

Let’s take a look at an example that demonstrates the use of the **new()** method:

```python
class MyMeta(type):
    def __new__(cls, name, bases, attrs):
        new_attrs = {}
        for attr_name, attr_value in attrs.items():
            if isinstance(attr_value, str):
                new_attrs[attr_name.upper()] = attr_value
            else:
                new_attrs[attr_name] = attr_value
        return super().__new__(cls, name, bases, new_attrs)
 
class MyClass(metaclass=MyMeta):
    x = 'hello'
    y = 123
    z = 'world'
 
print(MyClass.X)  # Output: "hello"
print(MyClass.Z)  # Output: "world"
# MyClass.y remains unchanged because its an integer
```

In this example, we define a new metaclass called MyMeta that modifies the class attributes by converting the values of any string attributes to uppercase. Then, we define a new class called MyClass that uses MyMeta as its metaclass. When we create an instance of MyClass, we can see that the string attributes have been converted to uppercase.

### init() Function

The init() method is responsible for initializing the class attributes. It takes three parameters:

* self: the newly created class object
* name: the name of the class being created
* bases: a tuple of base classes

Here’s an example of a custom metaclass that adds a new method to every class it creates:

``` python
class MyMeta(type):
    def __init__(self, name, bases):
        super().__init__(name, bases)
        self.foo = lambda self: print('Hello from foo!')
 
class MyClass(metaclass=MyMeta):
    pass
 
obj = MyClass()
obj.foo()  # Output: "Hello from foo!"
```

In this example, we define a new metaclass called MyMeta that adds a new method called foo() to every class it creates. Then, we define a new class called MyClass that uses MyMeta as its metaclass. When we create an instance of MyClass and call the foo() method, we can see that it works as expected.

In the next section of this tutorial, we will explore some more advanced use cases of a Metaclass in Python.

## Advanced Metaclass Features in Python

### Metaclass inheritance

Just like classes can inherit from other classes, metaclasses can also inherit from other metaclasses. This allows for the creation of complex hierarchies of metaclasses that can have their own special behaviors and attributes. To define a metaclass that inherits from another metaclass, simply specify the parent metaclass as the first argument when defining the new metaclass. For example:

``` python
class MyBaseMeta(type):
    pass

class MyChildMeta(MyBaseMeta):
    pass
```

In this example, we define a new metaclass called MyBaseMeta, and then define a child metaclass called MyChildMeta that inherits from MyBaseMeta.

Here’s a more detailed example of metaclass inheritance:

``` python
class BaseMeta(type):
    def __new__(cls, name, bases, attrs):
        print("Creating class", name, "with base classes", bases)
        return super().__new__(cls, name, bases, attrs)
 
class ChildMeta(BaseMeta):
    def __new__(cls, name, bases, attrs):
        print("Creating child class", name)
        return super().__new__(cls, name, bases, attrs)
 
class MyBaseClass(metaclass=BaseMeta):
    pass
 
class MyChildClass(metaclass=ChildMeta):
    pass
```

In this example, we define two metaclasses: BaseMeta and ChildMeta. ChildMeta inherits from BaseMeta.

When MyBaseClass and MyChildClass are defined, the **new** method of each respective metaclass is called. When MyBaseClass is defined, only BaseMeta‘s **new** method is called. However, when MyChildClass is defined, both ChildMeta‘s and BaseMeta‘s **new** methods are called, in that order.

Output from defining the MyBaseClass.

``` text
# Creating class MyBaseClass with base classes ()
```

Output from defining the MyChildClass.

``` text
# Creating class MyBaseClass with base classes ()
# Creating child class MyChildClass
# Creating class MyChildClass with base classes ()
```

## Multiple inheritance with metaclasses

Just like classes, metaclasses can also use multiple inheritance. This allows for even more complex hierarchies of metaclasses that can have their own unique behaviors and attributes. To specify multiple inheritance in a metaclass, simply provide a tuple of parent metaclasses as the first argument when defining the new metaclass. For example:

```python
class MyBaseMeta1(type):
    pass
 
class MyBaseMeta2(type):
    pass
 
class MyChildMeta(MyBaseMeta1, MyBaseMeta2):
    pass
```

In this example, we define two base metaclasses (MyBaseMeta1 and MyBaseMeta2), and then define a child metaclass called MyChildMeta that inherits from both base metaclasses.

## Dynamically generating classes with metaclasses

Metaclasses can be used to dynamically generate new classes at runtime. This can be useful in situations where you need to create many similar classes, or when you need to generate classes based on user input or configuration files. To generate a new class dynamically with a metaclass, simply call the metaclass with the appropriate arguments. For example:

```python
class MyMeta(type):
    pass

class MyClass1(metaclass=MyMeta):
    pass

class MyClass2(metaclass=MyMeta):
    pass

def create_class(class_name):
    return MyMeta(class_name, (), {})

MyClass3 = create_class('MyClass3')
```

In this example, we define a new metaclass called MyMeta, and then use it to generate three new classes (MyClass1, MyClass2, and MyClass3). The MyClass3 class is generated dynamically at runtime using the create_class() function, which calls the MyMeta metaclass with the appropriate arguments.

In the next section of this tutorial, we’ll take a look at some of the special attributes and methods that can be defined in a python metaclass.

## Python Metaclass: Special Attributes and Methods

In addition to the **\_\_new__**() and **\_\_init__**() methods, metaclasses can also define special attributes and methods that affect how new classes are created. Let’s take a look at some of the most commonly used metaclass attributes and methods.

### classcell()

This method is called when a new class is defined inside another class. It takes two arguments: the metaclass and the cell that contains the class definition.

``` python
class MyMeta(type):
    def classcell(cls, cell):
        print(f"Class '{cls.**name**}' defined in '{cell.cell_name}'")

class MyClass(metaclass=MyMeta):
    class MyNestedClass:
        pass
```

In this example, we define a new metaclass called MyMeta that defines the classcell() method. When we define a new nested class called MyNestedClass inside MyClass, the classcell() method is called with the MyMeta metaclass and the cell that contains the class definition. The method prints out a message indicating the name of the class and the name of the cell where it was defined.

### prepare()

This method is called before the \_\_new__() method when a new class is created. It takes three arguments: the metaclass, the name of the class, and the list of base classes. It returns a new dictionary of class attributes that will be used to create the new class.

```python
class MyMeta(type):
    def prepare(cls, name, bases):
        print(f"Preparing class '{name}'")
        return {'custom_attribute': 'Hello, world!'}

class MyClass(metaclass=MyMeta):
    pass

print(MyClass.custom_attribute)  # Output: "Hello, world!"
```

In this example, we define a new metaclass called MyMeta that defines the prepare() method. When we create a new class called MyClass with MyMeta as its metaclass, the prepare() method is called with MyMeta, the name 'MyClass', and an empty tuple for the base classes.

The method returns a new dictionary containing a single attribute, custom_attribute, with the value 'Hello, world!'. This dictionary is then used to create the new class object.

### \_\_instancecheck__() and \_\_subclasscheck__()

These methods are called when checking the type of an object or the subclass relationship between two classes. They take two arguments: the metaclass and the object or class being checked. They should return True if the object or class satisfies the type or subclass relationship, and False otherwise.

```python
class MyMeta(type):
    def __instancecheck__(cls, instance):
        print(f"Checking instance of '{cls.**name**}'")
        return isinstance(instance, str)

    def __subclasscheck__(cls, subclass):
        print(f"Checking subclass of '{cls.__name__}'")
        return issubclass(subclass, str)

class MyString(str, metaclass=MyMeta):
    pass

my_string = MyString('Hello, world!')

# Output: "Checking instance of 'MyString'", "True"

print(isinstance(my_string, MyString))

# Output: "Checking subclass of 'MyString'", "True"

print(issubclass(str, MyString))  
```

In this example, we define a new metaclass called MyMeta that defines the **\_\_instancecheck__**() and **\_\_subclasscheck__**() methods. When we define a new class called MyString that inherits from str and uses MyMeta as its metaclass, these methods are called when we check the instance type of my_string and the subclass relationship between str and MyString.

The **\_\_instancecheck__**() method prints out a message indicating that we are checking an instance of 'MyString', and returns True if the instance is a string. The **\_\_subclasscheck__**() method prints out a message indicating that we are checking a subclass of 'MyString', and returns True if the subclass is a string.
