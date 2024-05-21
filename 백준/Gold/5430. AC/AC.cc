#include <bits/stdc++.h>
#include <iostream>

using namespace std;

void parse(string& tmp, deque<int>& d){
  int cur = 0;
  for(int i = 1; i+1 < tmp.size(); i++)
  {
    if(tmp[i] == ','){
      d.push_back(cur);
      cur = 0;
    }
    else{
      cur = 10 * cur + (tmp[i] - '0');
    }
  }
  if(cur != 0)
    d.push_back(cur);
}

void print_result(deque<int>& d){
  cout << '[';
  for(int i = 0; i < d.size(); i++)
  {
    cout << d[i];
    if(i+1 != d.size())
      cout << ',';
  }
  cout << "]\n";
}

int main()
{
    int t,n;
    int flag=-1;
    string cmd,s;
    cin>>t;
    deque<int> dq;
 
    for(int i=0;i<t;i++){
        flag=-1;
        cin>>cmd;
        cin>>n;
        cin>>s;
        parse(s,dq);
        for(auto c:cmd){
            if(c=='R'){
                flag*=-1;
            }
            else{
                if(dq.empty()){
                    cout<<"error"<<"\n";
                    flag=2;
                    break;
                }
                if(flag<0){
                    dq.pop_front();
                }
                else{
                    dq.pop_back();
                }
            }
        }
        if(flag!=2){
            cout<<"[";
            if(flag==-1){
                
                for(int i=0;i<int(dq.size());i++){
                    if(i==int(dq.size())-1){
                        cout<<dq[i];
                    }
                    else{
                        cout<<dq[i]<<",";
                    }
                }
                
            }
            else if(flag==1){
                
                for(int i=int(dq.size())-1;i>=0;i--){
                    if(i==0){
                        cout<<dq[i];
                    }
                    else{
                        cout<<dq[i]<<",";
                    }
                }
                
            }
            cout<<"]\n";
            dq.clear();
        }
    }
    return 0;
}
