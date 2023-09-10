Steps
=====

To commemorate World Health Day, the school organized a campaign "Move!".
Every participant of the campaign had to register online,
specifying their team name and step length, and then input the number of steps taken each day.

Write a program that would provide a list of the total distance
covered by all students participating in the campaign for each team.

# Initial Data

Data is provided as **standard input** (keyboard) in the following format:

 * Lines contain the data of the participating students:
    * student team's name;
    * step length in **centimeters**;
    * the number of steps taken each day of the campaign.

If a student did not input the number of steps taken on a particular day, the value is 0.

Data is separated by a single space.

# Expected Results

Present the results as **standard output** (print to the screen):

* Output a list of the total distance covered by students in each team:
   * team's name;
   * the number of students who entered steps for **every** day of the campaign;
   * the total distance covered by the specified team in **kilometers**, rounded to two decimal places.

If a participant did not input the number of steps taken on at least one day, their data is not used in the calculations.

Output data is separated by a single space.

* One line should contain the data for one team.

If there wasn't a single student in any team name who entered the number of steps every day of the campaign,
then data for that team name does not need to be provided.
The list should be presented in the order as they were specified in the initial data file.

# Example Input

```
AAA 76 4353 8738 4946 6484 8946 6336 6284
BBB 74 4352 4578 7239 3157 7856 4648 5850
CCC 70 5731 9017 7641 6785 8865 3184 6800
DDD 68 9147 9702 8671 0 5398 9555 7141
EEE 71 7550 0 9089 8446 4390 0 7901
AAA 68 4227 7160 9565 3297 4401 5740 6908
AAA 73 6362 6705 0 9017 5971 3154 3585
CCC 73 7053 8752 7606 4079 5204 3819 3310
DDD 73 4263 5836 3336 5401 4719 0 0
EEE 71 8588 5617 6903 9565 6445 6786 4522
```

Where
 * Following lines (`AAA 76 4353 8738 4946 6484 8946 6336 6284`): team name, step length in _cm_, number of steps taken each day of the campaign
 * Value `0` means that the student did not enter the number of steps taken on that day

# Example Output

```
AAA 2 63.11
BBB 1 27.88
CCC 2 62.69
EEE 1 34.38
```

Where for each line:
 * First value (`AAA`): team name
 * Second value (`2`): number of students who entered the number of steps taken every day of the campaign
 * Third value (`63.11`): total distance covered in kilometers

# Limitations

- You cannot install any of the third party packages. You must rely only on your programming language of
choice standard library.
- All of the pre-installed software available in Github Actions [Ubuntu2204](https://github.com/actions/runner-images/blob/main/images/linux/Ubuntu2204-Readme.md) image is
available to use.
- Currently, we have set up examples for `js`, `php` and `go` programming languages. Feel free to add more!
