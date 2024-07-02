// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

library CCIPConfigArraysValidation {
    // errors
    error ArrayEmpty();
    error NotSorted(bytes32[] array);
    error HasDuplicates(bytes32[] array);
    error NotSubset(bytes32[] smaller, bytes32[] larger);

    function _checkSortedNoDuplicatesAndSubset(bytes32[] memory a, bytes32[] memory b) internal pure {
        if (a.length == 0 || b.length == 0) {
            revert ArrayEmpty();
        }

        _checkSortedAndDuplicates(a);
        _checkSortedAndDuplicates(b);

        // Check if 'a' is a subset of 'b'
        uint i = 0; // Pointer for 'a'
        uint j = 0; // Pointer for 'b'

        while (i < a.length && j < b.length) {
            if (a[i] > b[j]) {
                ++j; // Move the pointer in 'b' to find a match
            } else if (a[i] == b[j]) {
                ++i; // Found a match, move the pointer in 'a'
                ++j; // Also move in 'b' to continue checking
            } else {
                // 'a[i]' is less than 'b[j]' and no match is possible moving forward
                revert NotSubset(a, b);
            }
        }

        // If not all elements in 'a' were matched
        if (i < a.length) {
            revert NotSubset(a, b);
        }
    }

    // Helper function to check if array is sorted and has no duplicates
    function _checkSortedAndDuplicates(bytes32[] memory array) private pure {
        for (uint i = 1; i < array.length; ++i) {
            if (array[i] < array[i - 1]) {
                revert NotSorted(array);
            }
            if (array[i] == array[i - 1]) {
                revert HasDuplicates(array);
            }
        }
    }
}
