- Product (iGun): Defines the interface of objects the factory method creates.
- Concrete Product (ak47, maverick): Implements the Product interface.
- Creator (-): Derclares factory method, which returns an object of type Product. Creator may also define a default implementation of the factory method that returns a default Concrete Product object.; May call the factory method to create a  Product object.
- Concrete Creator (gunFactory.makeGun): Overrides the factory method to return an instance of a Concrete Product.

notes to myself:
- Concrete Products has parent class which is gun.
- No Creator, just the Concrete Creator. No default implementation (return error if the type doesn't match).