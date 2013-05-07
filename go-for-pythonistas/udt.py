

class Person(object):

    def __init__(self, firstname, lastname, phone):
        self.firstname = firstname
        self.lastname = lastname
        self.phone = phone

    def info(self):
        return "{} {} ({})".format(self.firstname, self.lastname, self.phone)


def main():
    p = Person("Dan", "Worth", "800-555-5555")
    print "Hello {}".format(p.info())

if __name__ == "__main__":
    main()
