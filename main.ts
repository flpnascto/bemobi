let initialAccount: number[];

function initialApplication(investmentsNumber: number): void {
  for (let i = 0; i < investmentsNumber; i++) {
    initialAccount.push(0);
  }
}

function applications(left: number, right: number, value: number): void {
  for (let i = left - 1; i < right; i++) {
    initialAccount[i] += value;
  }
}

const investment = [
  [1, 2, 10],
  [2, 4, 5],
  [3, 5, 12],
];

function(applicationsNumber: number, investments: number[][]): void {
  initialApplication(applicationsNumber);
  investments.forEach((investment) => {
    applications(investment[0], investment[1], investment[2]);
  });
}
