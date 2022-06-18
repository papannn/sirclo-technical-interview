#include<iostream>
#include<map>
#include<string>

using namespace std;

class Cart {
    private:
        map<string, int> inventories;
    
    public:
        int size() {
            return inventories.size();
        }

        int jumlahBuah(string kodeProduk) {
            if(inventories.find(kodeProduk) == inventories.end()) {
                throw "buah tidak ada dalam keranjang";
            }
            return inventories.at(kodeProduk);
        }

        void tambahProduk(string kodeProduk, int kuantitas) {
            if (kuantitas < 1) {
                throw "kuantitas tidak boleh kurang dari 1";
            }
            // If not exist
            if(inventories.find(kodeProduk) == inventories.end()) {
                inventories[kodeProduk] = kuantitas;
            } else {
                inventories[kodeProduk] += kuantitas;
            }
        }

        void tampilkanCart() {
            for(auto inventory : inventories) {
                cout << inventory.first << " (" << inventory.second << ")" << endl; 
            }
        }

        void hapusProduk(string kodeProduk) {
             // If exist
            if(inventories.find(kodeProduk) != inventories.end()) {
                inventories.erase(kodeProduk);
            }
        }
    
};
