Classes and Objects:
	- Class is a template for objects.
	- Object is an instance of a class.
	- Objects have attributes and methods
	- Class members are the variables and functions that belong to the class.
		Class MyClass {
			Public:
				Int myNum;
				String myString;
		};
	- Class keyword is used to create the class called MyClass
	- Public specifies access to be allowed from outside the object too.
		○ Private will specify that you can only access internally not even derived classes can access
		○ Protected is the middle zone that only the class or a derived class can access it
	- Can create an instance of a class without a constructor
		MyClass myObj;
	- Can create multiple instances of the class which are separate

Class Methods:
	- Methods can be made public, private and protected as well as variables.
	- First thing that really separates it from C structs.
	- Called with dot operator
		Class MyClass {
			Public:
				Void myMethod() { }
		};
		
		Int main() {
			MyClass myObj;
			myObj.myMethod();
			Return 0;
		}
	- Can also define outside of the class by using ::
		Void MyClass::myMethod() { }
		Int main () {
			MyClass myObj;
			myObj.myMethod();
			Return 0;
		}

Constructors:
	- Special method called upon creation of class automatically.
	- Named the same as the class
		Class MyClass {
			MyClass() {} 
		};
		Int main() {
			MyClass myObj;
			Return 0;
		}
	- Deconstructors also exist which are of the same name with a ~ in front
		Class MyClass {
			MyClass() {}
			~MyClass() {}
		};
		- Cannot have parameters nor can it have a return type.
		- Compiler will create a default constructor which calls delete/free
		- We want to write our own when we have dynamically allocated memory or pointers in the class.
		- Should free the memory before we destroy the instance otherwise we have a leak.
		- Need a virtual constructor in the base class if you have deconstructions in the derived class to ensure that the values 
	- Constructors are allowed parameters which can be used to set variables in the object normally.
	- Can have declaration and definition separately. So definition can be outside of the class where declaration is inside the class.

Access Specifiers: 
	- Public: members are accessible from outside the class
	- Private: members cannot be accessed (or viewed) from outside the class
	- Protected: members cannot be accessed from outside the class however, they can be accessed in inherited classes.
	- By default members of a class are private if no specifier is defined.

Encapsulation:
	- Concept of hiding sensitive data from users.
	- Values must be declared as private
	- Getter and setter methods must be provided to the user to allow for modification of the values present.
	- Provides increased security of the data and also allows the change of private data to be intentional.

Inheritance:
	- Methods and attributes can be inherited from one class to another.
	- Derived class (child) - class that inherits from another class
	- Base class (parent) - class being inherited from
	- To inherit from a class we simply use the ":" symbol
	```
	class Vehicle {
		public:
			string brand = "Ford";
			void honk() {
				cout << "Beep Beep \n";
			}
	};

	class Car: public Vehicle {
		public
			string model = "Mustang";
	};
	```
	- The derived class can use the honk method and also access the brand attribute.

Multilevel Inheritance:
	- Class can also be derived from one class which can already been derived from another class.

Multiple Inheritance:
	- class can be derived from more than one base class, using a comma separated list.
	```
	class MyClass {
		public:
			void myFunction() {}
	};

	class MyOtherClass {
		public:
			void myOtherFunction() {}
	};

	class MyChildClass: public MyClass, public MyOtherClass {};
	```

Inheritance Access:
	- private methods and attributes are unaccessible from the derived classes
	- public are available in the derived classes as well.
	- protected can be accessed within the inherited classes but not from outside a class instance.
	```
	class Employee {
		protected:
			int salary;
	};

	class Programmer: public Employee {
		public:
			int bonus;
			void setSalary(int s) {
				salary = s;
			}
			int getSalary() {
				return salary;
			}
	};
	```

Polymorphism:
	- occurs when we have many classes that are related to each other by inheritance.
	- Inheritance allows us to inherit attributes and methods from another class.
	- Polymorphism uses those methods to perform different tasks.
	- Allows the performing of single actions in many different ways.
	- I.e. redefine the behavior of a particular method.
	```
	class Animal {
		public:
			void animalSound() {}
	};
	class Pig : public Animal {
		public:
			void animalSound() { // different stuff }
	};
	class Dog : public Animal {
		public: 
			void animalSound() { // another different sound }
	};
	```
	- useful for code reusability.

Virtual Functions:
	- member function defined in base class and overridden in a derived class.
	- When referring to a pointer or reference to the base class you can call a virtual function for that object and it will execute the derived class's version of the function.
	- Rules:
		1. Cannot be static
		2. Can be friend functions of another class
		3. Should be accessed using pointer or reference of base class type to achieve runtime polymorphism
		4. prototype of virtual functions should be the same in base and derived classes.
		5. Not mandatory for derived class to override. Base class version used in that case.
		6. class may have a virtual destructor but cannot have virtual constructor.
	- Pure virtual function will have no implementation in the base class. 

Friend Functions / Classes:
	- A function that is declared outside of the object that can be used to access private and protected variables.
	- Must be indicated that a friend of this class. Signature must exist in the class declaration for the class with the friend keyword.
	```
	class Box {
		double width;
		public:
			friend void printWidth(Box box);
			void setWidth(double wid);
	}
	// member function
	void Box::setWidth(double wid) {
		width = wed;
	}
	// not a member but a friend
	void printWidth(Box box) {
		cout << box.width << endl;
	}
	```
	- More importantly a class can be made a friend
	- Allows only a couple classes to be able to access private data of a function
	- Protects the class privates from most classes with small exceptions.
	``` 
	class MyClass {
		friend class MyOtherClass;
		public:
			...
		private:
			...
	};
	```
	- This allows more fine grained control of classes access than the standard private, protected, public

Abstract Class:
	- Since c++ does not have the concept of interfaces a useful abstraction is abstract classes.
	- They will have no ability to be instantiated but can be used as a base class to inherit from.
	- Very popular in Golang and Java.
	- Class is made abstract by making at least one of its functions a pure virtual function
	- pure virtual made by
	```
	virtual int funcName() = 0;
	```
	- Can have any number of attributes as they will be uninitialized.
	```
	class Box {
		public:
			virtual double getVolume() = 0;
		private:
			double length;
			double width;
			double height;
	};
	```

Exceptions:
	- C++ allows for exceptions
	- Exceptions are thrown and handled by functions up the call stack.
	- try, catch blocks are used.
	- catch statements are used to catch exceptions which are thrown with the throw keyword.
	```
	try {
		// blah blah
		throw exception;
	}
	catch (ExceptionType ex) {
		// do something with it
	}
	```
	- can catch any type of exception by placing "..." in the catch variable list.

Generics:
	- allow type to be a parameter to methods, classes and interfaces.
	- great usage for classes that don't care what data they are holding such as data structures
	- Allows for large code reuse.
	- Advantages:
		1. Code Reusability
		2. Avoids function overloading
		3. write once use many times
	- Can be implemented using templates
	- Pass data as a parameter so we don't need to write same code for different types.
	- Generic Function Example:
	```
	template <typename T>
	T myMax (T x, T y) {
		return (x > y) ? x : y;
	}
	```
	- Generic Class Example:
	```
	template <typename T>
	class Array {
		private:
			T* ptr;
			int size;
		public:
			Array(T arr[], int s);
			void print();
	};

	template <typename T>
	Array<T>::Array(T arr[], int s) {
		prt = new T[s];
		size = s;
		for (int i = 0; i < size; i++) {
			ptr[i] = arr[i];
		}
	}

	template <typename T>
	void Array<T>::print() {
		for (int i = 0; i < size; ++i) {
			cout << " " << *(ptr + i);
		}
		cout << endl;
	}
	```

Solid Principles
	1. Single Responsibility Principle
		- Class should have one and only one reason to change. I.e. One job.
	2. Open Closed Principle
		- Objects should be open for extension for closed for modification. I.e. extensible without having to modify.
	3. Liskov Substitution Principle
		- Every base class should be substitutable for their base/parent class. I.e. when you call them there should be no big difference between which one gets called.
	4. Interface Segregation Principle
		- Client class should never be forced to implement an interface that it doesn't use, or depend on methods they don't use. I.e. no need to have methods that don't make sense just split into more interfaces.
	5. Dependency Inversion Principle
		- Entities must rely on abstractions, not on concretions. High-level module must not depend on the low-level module, but they should depend on abstractions. I.e. generic abstract classes should be provided for use in high level functions and objects. Allowing for easy swaps of say database engine rather than rewrites.
		- Allows for dependency injection.

Dependency Injection:
	- Transfer the task of creating the object to someone else and directly use the dependency.
	- DI acts as middle man which does all the work of creating the preferred implementation of an abstraction (object) providing it to the calling class.
	- Using SOLID you can be sure whatever implementation of the interface of the requested object will work. And allows for quick swapping.
	- Does have the potential to introduce run time errors as opposed to compile time. Which I do not prefer.


