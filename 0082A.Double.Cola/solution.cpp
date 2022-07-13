#include <iostream>
#include <vector>
#include <cmath>

using namespace std;

int main( void )
{
    // peopl's names
    vector< string > names;
    names.push_back( "Sheldon" );
    names.push_back( "Leonard" );
    names.push_back( "Penny" );
    names.push_back( "Rajesh" );
    names.push_back( "Howard" );

    // the nth can
    int n;
    cin >> n;

    int a = 0, s = names.size(), pa = 0;
    for ( int i = 1;; ++i )
    {
        a += s;
        if ( n > a )
        {
            pa = a;
            s += s;
            continue;
        }

        int ind = ( n - pa - 1 ) / static_cast< int >( pow( 2.0, i - 1 ) + 0.5 );

        // output the name of people who drinks the nth can
        cout << names[ ind ] << endl;
        break;
    }

    return 0;
}
