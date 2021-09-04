# dependency inversion principle

the dependency inversion principle is a specific form of loosely coupling software modules. When following this principle, the conventional dependency relationships established from high-level, policy-setting modules to low-level, dependency modules are reversed, thus rendering high-level modules independent of the low-level module implementation details.

The principle states:  
A. High-level modules should not import anything from low-level modules. Both should depend on abstractions (e.g., interfaces).  
B. Abstractions should not depend on details. Details (concrete implementations) should depend on abstractions.