# Axe Pool

CLI for managing and updating Bixaxe pools. Note that passwords are currently not implemented.

## Building The CLI

Will need Go installed, binary is `axe-pool`.

```
go build
```

## Help

Run without any arguments to see help.
```
./axe-pool
```

```
NAME:
   axe-pool - A cli for managing Bitaxe pools

USAGE:
   axe-pool [command] [options]

COMMANDS:
   list     List available pools and Bitaxe
   scan     Scan the network for Bitaxe
   add      Add a pool to the database
   del      Delete a pool from the database.  axe-pool del ID
   set      Update a bitaxe (or all) with primary and fallback example `axe-pool set all 3 2`
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## Usage

First scan the network for bitaxes. This creates a `axe-pool.db` sqlite database with basic information about the current Bitaxe pool settings.  Can be run periodically to update the database to add or remove Bitaxe.  Note the scanner assumes the network is /24 example: 192.168.1.0/24.

```
./axe-pool scan
```

## Add Some Pools

```
./axe-pool add
```

Will prompt for the pool host, port, and username. **TIP** leave off the worker name and axe-pool will use the Bitaxe's hostname when updating the pool settings.

## List Info

Use list to get pool and Bitaxe Ids. The primary and fallback pools are shown for each Bitaxe.

```
./axe-pool list

Pools:
ID |Host                 |User
1  |192.168.1.103:3333   |bc1qr4u3cdf3eur5mzhc2sh83sea7jykeucv9djre4
2  |solo.ckpool.org:3333 |bc1qr4u3cdf3eur5mzhc2sh83sea7jykeucv9djre4
3  |192.168.1.103:4333   |tb1qcy34mjamjgk322pjtmx95xqe2m5negcjkl8dyd
4  |dumb pool:3333       |test


Bitaxe:
ID |Host           |Stratum              |User
1  |ultra_204_led2 |192.168.1.103:4333   |tb1qcy34mjamjgk322pjtmx95xqe2m5negcjkl8dyd.ultra_204_led2
   |               |solo.ckpool.org:3333 |bc1qr4u3cdf3eur5mzhc2sh83sea7jykeucv9djre4.ultra_204_led2
   |               |                     |
2  |gamma_ss2_led5 |192.168.1.103:4333   |tb1qcy34mjamjgk322pjtmx95xqe2m5negcjkl8dyd.gamma_ss2_led5
   |               |solo.ckpool.org:3333 |bc1qr4u3cdf3eur5mzhc2sh83sea7jykeucv9djre4.gamma_ss2_led5
   |               |                     |
4  |gamma_bm_led3  |192.168.1.103:4333   |tb1qcy34mjamjgk322pjtmx95xqe2m5negcjkl8dyd.gamma_bm_led3
   |               |solo.ckpool.org:3333 |bc1qr4u3cdf3eur5mzhc2sh83sea7jykeucv9djre4.gamma_bm_led3
   |               |                     |
5  |gamma_ss_led4  |192.168.1.103:4333   |tb1qcy34mjamjgk322pjtmx95xqe2m5negcjkl8dyd.gamma_ss_led4
   |               |solo.ckpool.org:3333 |bc1qr4u3cdf3eur5mzhc2sh83sea7jykeucv9djre4.gamma_ss_led4
   |               |                     |
7  |Supra401_led1  |192.168.1.103:4333   |tb1qcy34mjamjgk322pjtmx95xqe2m5negcjkl8dyd.Supra401_led1
   |               |solo.ckpool.org:3333 |bc1qr4u3cdf3eur5mzhc2sh83sea7jykeucv9djre4.Supra401_led1
   |               |                     |
9  |ultra201_led6  |192.168.1.103:4333   |tb1qcy34mjamjgk322pjtmx95xqe2m5negcjkl8dyd.ultra201_led6
   |               |solo.ckpool.org:3333 |bc1qr4u3cdf3eur5mzhc2sh83sea7jykeucv9djre4.ultra201_led6
   |               |                     |
 ```  

## Finally Updating The Bitaxe

For setting a pool the arguments are `bitaxe_id` `primary_pool_id` `fallback_pool_id`. If the bitaxe_id is `all` then all bitaxes will be updated. **NOTE** This restarts the Bitaxe after updating the pool.

Single Bitaxe

```
./axe-pool set 2 1 2
```

Or for all bitaxes

```
./axe-pool set all 1 2
```

