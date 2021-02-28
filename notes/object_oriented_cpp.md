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
	- 
	


