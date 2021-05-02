#include<iostream>
using namespace std;
int main()
{
    int ln, ld, lh;
    cin>>ln>>ld>>lh;
    int d[ln];
    int h[ln];
    for (int i=0; i<ln; i++) {
        cin>>d[i]>>h[i];
    }
    double a = (double)lh/(double)ld;
    for (int i=0; i<ln; i++) {
        double ai = (double)(lh-h[i])/(double)(ld-d[i]);
        if (a > ai) {
            a = ai;
        }
    }
    double ans = ((double) lh) - a*((double) ld);
    cout << ans;
}