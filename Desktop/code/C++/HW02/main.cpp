#include <iostream>
#include "HW02.cpp"
using namespace std;
int main()
{
    Trip T;    
    T.insert_front("UENO");
    T.insert_front("KANDA");
    T.insert_front("TOKYO");
    T.insert_back("ASAKUSA");
    T.remove_back();
    T.printList();
    cout << T.size << endl;
    Station *lastStation;
    lastStation= T.visit(7,"RRLRRRL");
    cout<<lastStation->name;
    return 0;
}