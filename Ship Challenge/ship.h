#include<iostream>
#include<string>
#include<vector>

using namespace std;
class Ship {
    private:
        string kodeKapal;
    protected:
        string getKodeKapal() {
            return kodeKapal;
        }

        void setKodeKapal(string kodeKapal) {
            this->kodeKapal = kodeKapal;
        }

    public:
        virtual string anchored(string tujuan) = 0;
        virtual string sail(string tujuan) = 0;
};

class MotorBoat : public Ship {
    private:
        float persentaseBensin;
    public:
        MotorBoat(string kodeKapal, float persentaseBensin) {
            setKodeKapal(kodeKapal);
            this->persentaseBensin = persentaseBensin;
        }

        string anchored(string tujuan) {
            char buffer[200];
            sprintf(buffer, "Kapal Motor %s dengan jumlah bensin %f telah sampai di %s", getKodeKapal().c_str(), persentaseBensin, tujuan.c_str());
            return buffer;
        }

        string sail(string tujuan) {
            char buffer[200];
            sprintf(buffer, "Kapal Motor %s dengan jumlah bensin %f telah berlayar ke %s", getKodeKapal().c_str(), persentaseBensin, tujuan.c_str());
            return buffer;
        }
};

class SailShip : public Ship {
    private:
        int lebarLayar;
    public:
        SailShip(string kodeKapal, int lebarLayar) {
            setKodeKapal(kodeKapal);
            this->lebarLayar = lebarLayar;
        }

        string anchored(string tujuan) {
            char buffer[200];
            sprintf(buffer, "Kapal Layar %s dengan lebar layar %d cm telah sampai di %s", getKodeKapal().c_str(), lebarLayar, tujuan.c_str());
            return buffer;
        }

        string sail(string tujuan) {
            char buffer[200];
            sprintf(buffer, "Kapal Layar %s dengan lebar layar %d cm telah berlayar ke %s", getKodeKapal().c_str(), lebarLayar, tujuan.c_str());
            return buffer;
        }
};

class CruiseShip : public Ship {
    private:
        int jumlahRuangan;
    public:
        CruiseShip(string kodeKapal, int jumlahRuangan) {
            setKodeKapal(kodeKapal);
            this->jumlahRuangan = jumlahRuangan;
        }

        string anchored(string tujuan) {
            char buffer[200];
            sprintf(buffer, "Kapal Pesiar %s dengan jumlah ruangan %d kamar telah sampai di %s", getKodeKapal().c_str(), jumlahRuangan, tujuan.c_str());
            return buffer;
        }

        string sail(string tujuan) {
            char buffer[200];
            sprintf(buffer, "Kapal Pesiar %s dengan jumlah ruangan %d kamar telah berlayar ke %s", getKodeKapal().c_str(), jumlahRuangan, tujuan.c_str());
            return buffer;
        }
};

class Dock {
    private:
        vector<Ship*> ships;
    public:
        void addShip(Ship& ship) {
            ships.push_back(&ship);
        }

        void sailShip(string tujuan) {
            for(auto ship : ships) {
                cout << ship->sail(tujuan) << endl;
            }
            cout << endl;

            for(auto ship : ships) {
                cout << ship->anchored(tujuan) << endl;
            }
            cout << endl;
        }
};