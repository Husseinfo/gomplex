from argparse import ArgumentParser
from sys import argv

import grpc
from service_pb2 import ComplexNumber
from service_pb2_grpc import ComplexStub


def add(channel, numbers, sub=False):
    print(numbers)
    stub = ComplexStub(channel)
    complex_numbers = [ComplexNumber(real=complex(comp).real, imaginary=complex(comp).imag) for comp in numbers]
    print(f'{"Adding" if not sub else "subtracting"} numbers: {complex_numbers}')
    return stub.Add(iter(complex_numbers)) if not sub else stub.Sub(iter(complex_numbers))


def main():
    parser = ArgumentParser(description='Add two complex numbers - gRPC client')
    parser.add_argument('-H', '--host', type=str, default='localhost', help='Host address of service')
    parser.add_argument('-p', '--port', type=int, default=8000, help='TCP port of service')
    parser.add_argument('-n', '--number', action='append', type=str, help='A complex Number')
    parser.add_argument('-s', '--subtraction', action='store_true', default=False,
                        help='Performs subtraction instead of addition')
    if len(argv) < 2:
        parser.print_help()
        exit(1)
    args = parser.parse_args()
    channel = grpc.insecure_channel(f'{args.host}:{args.port}')
    print(f'Sum: {add(channel, args.number, args.subtraction)}')


if __name__ == '__main__':
    main()
