#include<iostream>
#include<string>
#include "cart.h"

class CartTest {
    private:
        static void Tambah_Produk_Kuantitas_Positif() {
            Cart keranjang;
            try {
                keranjang.tambahProduk("buah naga", 2);
                assert(true);
            } catch(const char* msg) {
                assert(false);
            }
        }

        static void Tambah_Produk_Kuantitas_Negatif() {
            Cart keranjang;
            try {
                keranjang.tambahProduk(string("buah naga"), -2);
                assert(false);
            } catch(const char* msg) {
                assert(true);
            }
        }

        static void Tambah_Keranjang_Cek_Ukuran() {
            Cart keranjang;
            try {
                keranjang.tambahProduk("buah naga", 2);
            } catch(const char* msg) {
                assert(false);
            }
            assert(keranjang.size() == 1);
        }

        static void Tambah_Keranjang_Cek_Ukuran_Banyak_Buah() {
            Cart keranjang;
            try {
                keranjang.tambahProduk("buah naga", 2);
                keranjang.tambahProduk("buah naga", 2);
                keranjang.tambahProduk("buah jambu", 2);
            } catch(const char* msg) {
                assert(false);
            }
            assert(keranjang.size() == 2);
        }

        static void Tambah_Keranjang_Cek_Ukuran_Terhapus() {
            Cart keranjang;
            try {
                keranjang.tambahProduk("buah naga", 2);
            } catch(const char* msg) {
                assert(false);
            }
            keranjang.hapusProduk("buah naga");
            assert(keranjang.size() == 0);
        }

        static void Tambah_Keranjang_Cek_Ukuran_Banyak_Buah_Terhapus() {
            Cart keranjang;
            try {
                keranjang.tambahProduk("buah naga", 2);
                keranjang.tambahProduk("buah naga", 2);
                keranjang.tambahProduk("buah jambu", 2);
            } catch(const char* msg) {
                assert(false);
            }
            keranjang.hapusProduk("buah naga");
            assert(keranjang.size() == 1);
        }

        static void Keranjang_Kosong_Hapus_Buah() {
            Cart keranjang;
            keranjang.hapusProduk("buah naga");
            assert(keranjang.size() == 0);
        }

        static void Keranjang_Buah_Cek_Kuantitas () {
            Cart keranjang;
            try {
                keranjang.tambahProduk("buah naga", 2);
            } catch(const char* msg) {
                assert(false);
            }
            assert(keranjang.jumlahBuah("buah naga") == 2);
        }

        static void Keranjang_Buah_Cek_Kuantitas_Multiple_Add() {
            Cart keranjang;
            try {
                keranjang.tambahProduk("buah naga", 2);
                keranjang.tambahProduk("buah naga", 3);
            } catch(const char* msg) {
                assert(false);
            }
            assert(keranjang.jumlahBuah("buah naga") == 5);
        }

        static void Keranjang_Buah_Cek_Kuantitas_Buah_Kosong() {
            Cart keranjang;
            try {
                keranjang.tambahProduk("buah naga", 2);
                keranjang.tambahProduk("buah naga", 3);
            } catch(const char* msg) {
                assert(false);
            }

            try {
                keranjang.jumlahBuah("buah pepaya");
                assert(false);
            } catch(const char* msg) {
                assert(true);
            }
        }


    
    public:
        static void RunTest() {
            Tambah_Produk_Kuantitas_Positif();
            Tambah_Produk_Kuantitas_Negatif();
            Tambah_Keranjang_Cek_Ukuran();
            Tambah_Keranjang_Cek_Ukuran_Banyak_Buah();
            Tambah_Keranjang_Cek_Ukuran_Terhapus();
            Tambah_Keranjang_Cek_Ukuran_Banyak_Buah_Terhapus();
            Keranjang_Kosong_Hapus_Buah();
            Keranjang_Buah_Cek_Kuantitas();
            Keranjang_Buah_Cek_Kuantitas_Multiple_Add();
            Keranjang_Buah_Cek_Kuantitas_Buah_Kosong();

            cout << "Unit Test Success!!!" << endl;
        }
};

int main() {
    CartTest::RunTest();
    return 0;
}