
#include <iostream>

using namespace std;

class Employee {

private: // Private attribute
    int salary;

public:
    // Setter
    void setSalary(int s) {
        
        salary = s;

    }

    // Getter
    int getSalary() {
       return(salary);
    }
};
