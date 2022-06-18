#include "ship.h"

int main() {
    Dock dock;
    MotorBoat kapalMotor("KPL-MTR999", 90);
    SailShip kapalLayar("KPL-LYR121", 200);
    CruiseShip kapalPesiar("KPL-PSR222", 4);

    dock.addShip(kapalMotor);
    dock.addShip(kapalLayar);
    dock.addShip(kapalPesiar);

    dock.sailShip("Bandung");
    dock.sailShip("Padang");
}