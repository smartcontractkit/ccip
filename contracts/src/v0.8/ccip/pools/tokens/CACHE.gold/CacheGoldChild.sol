// SPDX-License-Identifier: MIT
pragma solidity ^0.8.11;

import {IFxERC20} from "./IFxERC20.sol";
import "./lib/AccessControl.sol";

/// @title The CacheGold Token Contract for L2/Sidechains
/// @author CACHE TEAM
contract CacheGoldChild is IFxERC20, AccessControl {
    bytes32 public constant FEE_ENFORCER_ROLE = keccak256("FEE_ENFORCER_ROLE");
    // 10^8 shortcut
    uint256 private constant TOKEN = 10**8;

    string public constant name = "CACHE Gold";
    string public constant symbol = "CGT";
    uint8 public constant decimals = 8;

    // Seconds in a day
    uint256 private constant DAY = 86400;

    // Days in a year
    uint256 private constant YEAR = 365;

    // The maximum transfer fee is 10 basis points
    uint256 private constant MAX_TRANSFER_FEE_BASIS_POINTS = 10;

    // Basis points means divide by 10,000 to get decimal
    uint256 private constant BASIS_POINTS_MULTIPLIER = 10000;

    // The storage fee of 0.25%
    uint256 private constant STORAGE_FEE_DENOMINATOR = 4e10;

    // The inactive fee of 0.50%
    uint256 private constant INACTIVE_FEE_DENOMINATOR = 2e10;

    // The minimum balance that would accrue a storage fee after 1 day
    uint256 private constant MIN_BALANCE_FOR_FEES = 146000;

    // Initial basis points for transfer fee
    uint256 private _transferFeeBasisPoints = 10;

    // How many days need to pass before late fees can be collected (3 years)
    uint256 public constant INACTIVE_THRESHOLD_DAYS = 1095;

    // Token balance of each address
    mapping(address => uint256) private _balances;

    // Allowed transfer from address
    mapping(address => mapping(address => uint256)) private _allowances;

    // Last time storage fee was paid
    mapping(address => uint256) private _timeStorageFeePaid;

    // Last time the address produced a transaction on this contract
    mapping(address => uint256) private _timeLastActivity;

    // Amount of inactive fees already paid
    mapping(address => uint256) private _inactiveFeePaid;

    // If address doesn't have any activity for INACTIVE_THRESHOLD_DAYS
    // we can start deducting chunks off the address so that
    // full balance can be recouped after 200 years. This is likely
    // to happen if the user loses their private key.
    mapping(address => uint256) private _inactiveFeePerYear;

    // Addresses not subject to transfer fees
    mapping(address => bool) private _transferFeeExempt;

    // Address is not subject to storage fees
    mapping(address => bool) private _storageFeeExempt;

    // Save grace period on storage fees for an address
    mapping(address => uint256) private _storageFeeGracePeriod;
    // Current total number of tokens created
    uint256 private _totalSupply;

    // Address where storage and transfer fees are collected
    address private _feeAddress;

    // A fee-exempt address that can be used to collect gold tokens in exchange
    // for redemption of physical gold
    address private _redeemAddress;

    //addresses related to the fxManager
    address internal _fxManager;
    address internal _connectedToken;


    // Grace period before storage fees kick in
    uint256 private _storageFeeGracePeriodDays = 0;

    // When gold bars are minted on child chain
    event Mint(uint256 amount, address user);

    // When a user burns tokens in child for withdrawal to mainnet
    event withdrawBurn(address _from, uint256 _amount);

    // When an account has no activity for INACTIVE_THRESHOLD_DAYS
    // it will be flagged as inactive
    event AccountInactive(address indexed account, uint256 feePerYear);

    // If an previoulsy dormant account is reactivated
    event AccountReActive(address indexed account);

    // Emit if the Operator address is changed
    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    // Emit if the critical addresses are changed
    event AddressChange(string addressType, address indexed account);

    // Emit if a critical fee is changed
    event FeeChange(string feeType, uint fee);

    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        setFeeExempt(msg.sender);
    }

    function initialize(
        address __feeAddress,
        address __feeEnforcer,
        address __fxManager_,
        address __connectedToken,
        address __redeemAddress
    ) external onlyRole(DEFAULT_ADMIN_ROLE) override  {
        require(__fxManager_ != address(0x0) && __connectedToken != address(0x0), "Zero address inputted");
        require(_fxManager == address(0x0) && _connectedToken == address(0x0), "Token is already initialized");
        _fxManager = __fxManager_;
        _connectedToken = __connectedToken;
        _redeemAddress = __redeemAddress;
        _feeAddress = __feeAddress;
        _grantRole(FEE_ENFORCER_ROLE, __feeEnforcer);
        setFeeExempt(_feeAddress);
        setFeeExempt(_fxManager);
    }

    // fxManager returns fx manager
    function fxManager() external view override returns (address) {
        return _fxManager;
    }

    // connectedToken returns root token
    function connectedToken() external view override returns (address) {
        return _connectedToken;
    }

    function setFxManager(address __fxManager) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _fxManager = __fxManager;
        setFeeExempt(_fxManager);
    }


    /**
     * @dev Transfer token for a specified address
     * @param to The address to transfer to.
     * @param value The amount to be transferred.
     */
    function transfer(address to, uint256 value)
    external
    override
    returns (bool)
    {
        require(_balances[msg.sender] >= value, "Insufficient Balance To Make This Transfer");

        // Update activity for the sender
        _updateActivity(msg.sender);

        // Can opportunistically mark an account inactive if someone
        // sends money to it
        if (_shouldMarkInactive(to)) {
            _setInactive(to);
        }

        _transfer(msg.sender, to, value);
        return true;
    }

    /**
     * @dev Approve the passed address to spend the specified amount of tokens on behalf of msg.sender.
     * Beware that changing an allowance with this method brings the risk that someone may use both the old
     * and the new allowance by unfortunate transaction ordering. One possible solution to mitigate this
     * race condition is to first reduce the spender's allowance to 0 and set the desired value afterwards:
     * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
     * @param spender The address which will spend the funds.
     * @param value The amount of tokens to be spent.
     */
    function approve(address spender, uint256 value)
    external
    override
    returns (bool)
    {
        _updateActivity(msg.sender);
        _approve(msg.sender, spender, value);
        return true;
    }

    /**
     * @dev Transfer tokens from one address to another.
     * Note that while this function emits an Approval event, this is not required as per the specification,
     * and other compliant implementations may not emit the event.
     * Also note that even though balance requirements are not explicitly checked,
     * any transfer attempt over the approved amount will automatically fail due to
     * SafeMath revert when trying to subtract approval to a negative balance
     * @param from address The address which you want to send tokens from
     * @param to address The address which you want to transfer to
     * @param value uint256 the amount of tokens to be transferred
     */
    function transferFrom(
        address from,
        address to,
        uint256 value
    ) external override returns (bool) {
        _updateActivity(msg.sender);
        _transfer(from, to, value);
        _approve(from, msg.sender, (_allowances[from][msg.sender] - (value)));
        return true;
    }

    /**
     * @dev Increase the amount of tokens that an owner allowed to a spender.
     * approve should be called when allowed_[_spender] == 0. To increment
     * allowed value is better to use this function to avoid 2 calls (and wait until
     * the first transaction is mined)
     * From MonolithDAO Token.sol
     * Emits an Approval event.
     * @param spender The address which will spend the funds.
     * @param addedValue The amount of tokens to increase the allowance by.
     */
    function increaseAllowance(address spender, uint256 addedValue)
    external
    returns (bool)
    {
        _updateActivity(msg.sender);
        _approve(
            msg.sender,
            spender,
            (_allowances[msg.sender][spender] + (addedValue))
        );
        return true;
    }

    /**
     * @dev Decrease the amount of tokens that an owner allowed to a spender.
     * approve should be called when allowed_[_spender] == 0. To decrement
     * allowed value is better to use this function to avoid 2 calls (and wait until
     * the first transaction is mined)
     * From MonolithDAO Token.sol
     * Emits an Approval event.
     * @param spender The address which will spend the funds.
     * @param subtractedValue The amount of tokens to decrease the allowance by.
     */
    function decreaseAllowance(address spender, uint256 subtractedValue)
    external
    returns (bool)
    {
        _updateActivity(msg.sender);
        _approve(
            msg.sender,
            spender,
            (_allowances[msg.sender][spender] - (subtractedValue))
        );
        return true;
    }

    /**
     * @dev Manually pay storage fees on senders address. Exchanges may want to
     * periodically call this function to pay owed storage fees. This is a
     * cheaper option than 'send to self', which would also trigger paying
     * storage fees
     *
     * @return A boolean that indicates if the operation was successful.
     */
    function payStorageFee() external returns (bool) {
        _updateActivity(msg.sender);
        _payStorageFee(msg.sender);
        return true;
    }

    function setAccountInactive(address account)
    external
    onlyRole(FEE_ENFORCER_ROLE)
    returns (bool)
    {
        require(
            _shouldMarkInactive(account),
            "Account not eligible to be marked inactive"
        );
        _setInactive(account);
        return true;
    }

    /**
     * @dev Contract allows the forcible collection of storage fees on an address
     * if it is has been more than than 365 days since the last time storage fees
     * were paid on this address.
     *
     * Alternatively inactive fees may also be collected periodically on a prorated
     * basis if the account is currently marked as inactive.
     *
     * @param account The address to pay storage fees on
     * @return A boolean that indicates if the operation was successful.
     */
    function forcePayFees(address account)
    external
    onlyRole(FEE_ENFORCER_ROLE)
    returns (bool)
    {
        require(account != address(0), "Zero address used");
        require(
            _balances[account] > 0,
            "Account has no balance, cannot force paying fees"
        );

        // If account is inactive, pay inactive fees
        if (isInactive(account)) {
            uint256 paid = _payInactiveFee(account);
            require(paid > 0, "Error no fees paid!");
        } else if (_shouldMarkInactive(account)) {
            // If it meets inactive threshold, but hasn't been set yet, set it.
            // This will also trigger automatic payment of owed storage fees
            // before starting inactive fees
            _setInactive(account);
        } else {
            // Otherwise just force paying owed storage fees, which can only
            // be called if they are more than 365 days overdue
            require(
                daysSincePaidStorageFee(account) >= YEAR,
                "Account has paid storage fees more recently than 365 days"
            );
            uint256 paid = _payStorageFee(account);
            require(
                paid > 0,
                "No appreciable storage fees due, will refund gas"
            );
        }
        return true;
    }

    /**
     * @dev Set the address to collect fees
     * @param newFeeAddress The address to collect storage and transfer fees
     * @return An bool representing successfully changing fee address
     */
    function setFeeAddress(address newFeeAddress)
    external
    onlyRole(DEFAULT_ADMIN_ROLE)
    returns (bool)
    {
        require(newFeeAddress != address(0), "Zero address used");
        _feeAddress = newFeeAddress;
        setFeeExempt(_feeAddress);
        emit AddressChange("Fee Address", newFeeAddress);
        return true;
    }

    /**
    * @dev Set the address to deposit tokens when redeeming for physical locked bars.
    * @param newRedeemAddress The address to redeem tokens for bars
    * @return An bool representing successfully changing redeem address
    */
    function setRedeemAddress(address newRedeemAddress) external onlyRole(DEFAULT_ADMIN_ROLE) returns(bool) {
        require(newRedeemAddress != address(0), "Zero address used");
        _redeemAddress = newRedeemAddress;
        setFeeExempt(_redeemAddress);
        emit AddressChange("Redeem Address", newRedeemAddress);
        return true;
    }

    /**
     * @dev Set the number of days before storage fees begin accruing.
     * @param daysGracePeriod The global setting for the grace period before storage
     * fees begin accruing. Note that calling this will not change the grace period
     * for addresses already actively inside a grace period
     */
    function setStorageFeeGracePeriodDays(uint256 daysGracePeriod)
    external
    onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _storageFeeGracePeriodDays = daysGracePeriod;
        emit FeeChange("Storage Fee Grace Period Days", daysGracePeriod);
    }

    /**
     * @dev Set this account as being exempt from transfer fees. This may be used
     * in special circumstance for cold storage addresses owed by Cache, exchanges, etc.
     * @param account The account to exempt from transfer fees
     */
    function setTransferFeeExempt(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _transferFeeExempt[account] = true;
    }

    /**
     * @dev Set this account as being exempt from storage fees. This may be used
     * in special circumstance for cold storage addresses owed by Cache, exchanges, etc.
     * @param account The account to exempt from storage fees
     */
    function setStorageFeeExempt(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _storageFeeExempt[account] = true;
    }

    /**
     * @dev Set a new transfer fee in basis points, must be less than or equal to 10 basis points
     * @param fee The new transfer fee in basis points
     */
    function setTransferFeeBasisPoints(uint256 fee) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(
            fee <= MAX_TRANSFER_FEE_BASIS_POINTS,
            "Transfer fee basis points must be an integer between 0 and 10"
        );
        emit FeeChange("Transfer Fee Basis Points", fee);
        _transferFeeBasisPoints = fee;
    }

    /**
     * @dev Gets the balance of the specified address deducting owed fees and
     * accounting for the maximum amount that could be sent including transfer fee
     * @param owner The address to query the balance of.
     * @return An uint256 representing the amount sendable by the passed address
     * including transaction and storage fees
     */
    function balanceOf(address owner) external view override returns (uint256) {
        return maximumTransferAmount(owner);
    }

    /**
     * @dev Gets the balance of the specified address not deducting owed fees.
     * this returns the 'traditional' ERC-20 balance that represents the balance
     * currently stored in contract storage.
     * @param owner The address to query the balance of.
     * @return An uint256 representing the amount stored in passed address
     */
    function balanceOfNoFees(address owner) external view returns (uint256) {
        return _balances[owner];
    }

    /**
    * @return address for redeeming tokens for gold bars
    */
    function redeemAddress() external view returns(address) {
        return _redeemAddress;
    }

    /**
     * @return address where fees are collected
     */
    function getFeeAddress() external view returns (address) {
        return _feeAddress;
    }

    /**
     * @return the current number of days and address is exempt
     * from storage fees upon receiving tokens
     */
    function storageFeeGracePeriodDays() external view returns (uint256) {
        return _storageFeeGracePeriodDays;
    }

    /**
     * @return the current transfer fee in basis points [0-10]
     */
    function transferFeeBasisPoints() external view returns (uint256) {
        return _transferFeeBasisPoints;
    }

    /**
     * @dev Function to mint certain amount of tokens in the child chain
     * @param user The user to whom the minted tokens have to be sent to
     * @param amount The amount of tokens to add to the supply and pass to the user
     */
    function mint(address user, uint256 amount)
    external
    override
    {
        require(msg.sender == _fxManager, "Invalid sender");
        _totalSupply = _totalSupply + amount;
        uint storageFeeTo = calcStorageFee(user);// automatically deduct any pending storage fee
        _balances[user] = _balances[user] + amount - storageFeeTo;
        emit Mint(amount,user);
        if(_timeStorageFeePaid[user] == 0){
            //checks if it is the first time the user is depositing gold into the account
            _storageFeeGracePeriod[user] = _storageFeeGracePeriodDays;
            _timeLastActivity[user] = block.timestamp;
            _timeStorageFeePaid[user] = block.timestamp;
        }
    }

    /**
     * @dev Function to burn certain amount of tokens in the child chain
     * @param account The account from whom the tokens have to be removed from
     * @param amount The amount of tokens to remove from the supply and burn
     */
    function burn(address account, uint256 amount)
    external
    override
    {
        require(msg.sender == _fxManager, "Invalid sender");
        uint256 currentAllowance = allowance(account, msg.sender);
        require(
            currentAllowance >= amount,
            "ERC20: burn amount exceeds allowance"
        );
        _approve(account, msg.sender, currentAllowance - amount);
    unchecked
    {
        _balances[account] = _balances[account] - amount;
    }

        _totalSupply = _totalSupply - amount;//reduce the total supply
        emit withdrawBurn(account, amount);
        emit Transfer(account, address(0), amount);// Fx Tunnel expects an event denoting a burn to withdraw on mainnet
    }

    function totalCirculation() external view returns (uint256) {
        return _totalSupply;
    }

    function totalSupply() external view override returns (uint256) {
        return _totalSupply;
    }

    /**
     * @dev Function to check the amount of tokens that an owner allowed to a spender.
     * @param owner address The address which owns the funds.
     * @param spender address The address which will spend the funds.
     * @return A uint256 specifying the amount of tokens still available for the spender.
     */
    function allowance(address owner, address spender)
    public
    view
    override
    returns (uint256)
    {
        return _allowances[owner][spender];
    }

    /**
     * @dev Set account is no longer exempt from all fees
     * @param account The account to reactivate fees
     */
    function unsetFeeExempt(address account) public onlyRole(DEFAULT_ADMIN_ROLE) {
        _transferFeeExempt[account] = false;
        _storageFeeExempt[account] = false;
    }
    /**
     * @dev Set this account as being exempt from all fees. This may be used
     * in special circumstance for cold storage addresses owed by Cache, exchanges, etc.
     * @param account The account to exempt from storage and transfer fees
     */
    function setFeeExempt(address account) public onlyRole(DEFAULT_ADMIN_ROLE) {
        _transferFeeExempt[account] = true;
        _storageFeeExempt[account] = true;
    }

    /**
     * @dev Check if the address given is exempt from storage fees
     * @param account The address to check
     * @return A boolean if the address passed is exempt from storage fees
     */
    function isStorageFeeExempt(address account) public view returns (bool) {
        return _storageFeeExempt[account];
    }

    /**
     * @dev Check if the address given is exempt from transfer fees
     * @param account The address to check
     * @return A boolean if the address passed is exempt from transfer fees
     */
    function isTransferFeeExempt(address account) public view returns (bool) {
        return _transferFeeExempt[account];
    }

    /**
     * @dev Check if the address given is exempt from transfer fees
     * @param account The address to check
     * @return A boolean if the address passed is exempt from transfer fees
     */
    function isAllFeeExempt(address account) public view returns (bool) {
        return _transferFeeExempt[account] && _storageFeeExempt[account];
    }


    function isInactive(address account) public view returns (bool) {
        return _inactiveFeePerYear[account] > 0;
    }

    /**
     * @dev Get the number of days since the account last paid storage fees
     * @param account The address to check
     * @return A uint256 representing the number of days since storage fees where last paid
     */
    function daysSincePaidStorageFee(address account)
    public
    view
    returns (uint256)
    {
        if (isInactive(account) || _timeStorageFeePaid[account] == 0) {
            return 0;
        }
        return (block.timestamp - _timeStorageFeePaid[account]) / (DAY);
    }

    /**
     * @dev Get the days since the account last sent a transaction to the contract (activity)
     * @param account The address to check
     * @return A uint256 representing the number of days since the address last had activity
     * with the contract
     */
    function daysSinceActivity(address account) public view returns (uint256) {
        if (_timeLastActivity[account] == 0) {
            return 0;
        }
        return (block.timestamp - (_timeLastActivity[account])) / (DAY);
    }

    /**
     * @dev Returns the total number of fees owed on a particular address
     * @param account The address to check
     * @return The total storage and inactive fees owed on the address
     */
    function calcOwedFees(address account) public view returns (uint256) {
        return calcStorageFee(account) + (calcInactiveFee(account));
    }

    /**
     * @dev Calculate the current storage fee owed for a given address
     * @param account The address to check
     * @return A uint256 representing current storage fees for the address
     */
    function calcStorageFee(address account) public view returns (uint256) {
        // If an account is in an inactive state those fees take over and
        // storage fees are effectively paused
        uint256 balance = _balances[account];
        if (
            isInactive(account) || isStorageFeeExempt(account) || balance == 0
        ) {
            return 0;
        }

        uint256 daysSinceStoragePaid = daysSincePaidStorageFee(account);
        uint256 daysInactive = daysSinceActivity(account);
        uint256 gracePeriod = _storageFeeGracePeriod[account];

        // If there is a grace period, we can deduct it from the daysSinceStoragePaid
        if (gracePeriod > 0) {
            if (daysSinceStoragePaid > gracePeriod) {
                daysSinceStoragePaid = daysSinceStoragePaid - (gracePeriod);
            } else {
                daysSinceStoragePaid = 0;
            }
        }

        if (daysSinceStoragePaid == 0) {
            return 0;
        }

        // This is an edge case where the account has not yet been marked inactive, but
        // will be marked inactive whenever there is a transaction allowing it to be marked.
        // Therefore we know storage fees will only be valid up to a point, and inactive
        // fees will take over.
        if (daysInactive >= INACTIVE_THRESHOLD_DAYS) {
            // This should not be at risk of being negative, because its impossible to force paying
            // storage fees without also setting the account to inactive, so if we are here it means
            // the last time storage fees were paid was BEFORE the account became eligible to be inactive
            // and it's always the case that daysSinceStoragePaid > daysInactive - (INACTIVE_THRESHOLD_DAYS)
            daysSinceStoragePaid =
            daysSinceStoragePaid -
            (daysInactive - (INACTIVE_THRESHOLD_DAYS));
        }
        // The normal case with normal storage fees
        return storageFee(balance, daysSinceStoragePaid);
    }

    /**
     * @dev Calculate the current inactive fee for a given address
     * @param account The address to check
     * @return A uint256 representing current inactive fees for the address
     */
    function calcInactiveFee(address account) public view returns (uint256) {
        uint256 balance = _balances[account];
        uint256 daysInactive = daysSinceActivity(account);

        // if the account is marked inactive already, can use the snapshot balance
        if (isInactive(account)) {
            return
            _calcInactiveFee(
                balance,
                daysInactive,
                _inactiveFeePerYear[account],
                _inactiveFeePaid[account]
            );
        } else if (_shouldMarkInactive(account)) {
            // Account has not yet been marked inactive in contract, but the inactive fees will still be due.
            // Just assume snapshotBalance will be current balance after fees
            uint256 snapshotBalance = balance - (calcStorageFee(account));
            return
            _calcInactiveFee(
                snapshotBalance, // current balance
                daysInactive, // number of days inactive
                _calcInactiveFeePerYear(snapshotBalance), // the inactive fee per year based on balance
                0
            ); // fees paid already
        }
        return 0;
    }

    /**
     * @dev Calculate the amount that would clear the balance from the address
     * accounting for owed storage and transfer fees
     * accounting for storage and transfer fees
     * @param account The address to check
     * @return A uint256 representing total amount an address has available to send
     */
    function maximumTransferAmount(address account) public view returns (uint256) {
        require(account != address(0), "Zero address used");

        // Internal addresses pay no fees, so they can send their entire balance
        uint256 balanceAfterStorage = _balances[account] -
        (calcOwedFees(account));
        if (_transferFeeBasisPoints == 0 || isTransferFeeExempt(account)) {
            return balanceAfterStorage;
        }

        // Edge cases where remaining balance is 0.00000001, but is effectively 0
        if (balanceAfterStorage <= 1) {
            return 0;
        }

        // Calculate the send all amount including storage fee
        // Send All = Balance / 1.001
        // and round up 0.00000001
        uint256 divisor = TOKEN +
        (_transferFeeBasisPoints * BASIS_POINTS_MULTIPLIER);
        uint256 sendAllAmount = (((balanceAfterStorage * TOKEN) /
        divisor) +
        1);

        // Calc transfer fee on send all amount
        uint256 transFee = (sendAllAmount * (_transferFeeBasisPoints)) /
        (BASIS_POINTS_MULTIPLIER);

        // Fix to include rounding errors
        if ((sendAllAmount + transFee) > balanceAfterStorage) {
            return sendAllAmount - (1);
        }

        return sendAllAmount;
    }

    /*
     * @dev Calculate the transfer fee on an amount
     * @param value The value being sent
     * @return A uint256 representing the transfer fee on sending the value given
     */
    function calcTransferFee(address account, uint256 value)
    public
    view
    returns (uint256)
    {
        if (isTransferFeeExempt(account)) {
            return 0;
        }
        // Basis points -> decimal multiplier:
        // f(x) = x / 10,0000 (10 basis points is 0.001)
        // So transfer fee working with integers =
        // f(balance, basis) = (balance * TOKEN) / (10,000 * TOKEN / basis)
        return (value * (_transferFeeBasisPoints)) / (BASIS_POINTS_MULTIPLIER);
    }

    /*
     * @dev Calculate the storage fee for a given balance after a certain number of
     * days have passed since the last time fees were paid.
     * @param balance The current balance of the address
     * @param daysSinceStoragePaid The number days that have passed since fees where last paid
     * @return A uint256 representing the storage fee owed
     */
    function storageFee(uint256 balance, uint256 daysSinceStoragePaid)
    public
    pure
    returns (uint256)
    {
        uint256 fee = (balance * TOKEN * daysSinceStoragePaid /
        YEAR) /
        STORAGE_FEE_DENOMINATOR;
        if (fee > balance) {
            return balance;
        }
        return fee;
    }

    /**
     * @dev Approve an address to spend another addresses' tokens.
     * @param owner The address that owns the tokens.
     * @param spender The address that will spend the tokens.
     * @param value The number of tokens that can be spent.
     */
    function _approve(
        address owner,
        address spender,
        uint256 value
    ) internal {
        require(spender != address(0), "Zero address used");
        require(owner != address(0), "Zero address used");

        _allowances[owner][spender] = value;
        emit Approval(owner, spender, value);
    }

    /**
     * @dev Transfer token for a specified addresses. Transfer is modified from a
     * standard ERC20 contract in that it must also process transfer and storage fees
     * for the token itself. Additionally there are certain internal addresses that
     * are not subject to fees.
     * @param from The address to transfer from.
     * @param to The address to transfer to.
     * @param value The amount to be transferred.
     */
    function _transfer(
        address from,
        address to,
        uint256 value
    ) internal {
        require(from != address(0), "Zero address used");
        require(to != address(0), "Zero address used");
        require(to != address(this), "Cannot transfer tokens to the contract");

        // redeem address can only call burn
        require(from != _redeemAddress,
            "Redeem address can only transfer to mainnet by burning");
        // If the account was previously inactive and initiated the transfer, the
        // inactive fees and storage fees have already been paid by the time we get here
        // via the _updateActivity() call
        uint256 storageFeeFrom = calcStorageFee(from);
        uint256 storageFeeTo = 0;
        uint256 allFeeFrom = storageFeeFrom;
        uint256 balanceFromBefore = _balances[from];
        uint256 balanceToBefore = _balances[to];

        // If not sending to self can pay storage and transfer fee
        if (from != to) {
            // Need transfer fee and storage fee for receiver if not sending to self
            allFeeFrom = allFeeFrom + (calcTransferFee(from, value));
            storageFeeTo = calcStorageFee(to);
            _balances[from] = balanceFromBefore - (value) - (allFeeFrom);
            _balances[to] = balanceToBefore + (value) - (storageFeeTo);
            _balances[_feeAddress] =
            _balances[_feeAddress] +
            (allFeeFrom) +
            (storageFeeTo);
        } else {
            // Only storage fee if sending to self
            _balances[from] = balanceFromBefore - (storageFeeFrom);
            _balances[_feeAddress] = _balances[_feeAddress] + (storageFeeFrom);
        }

        // Regular Transfer
        emit Transfer(from, to, value);

        // Fee transfer on `from` address
        if (allFeeFrom > 0) {
            emit Transfer(from, _feeAddress, allFeeFrom);
            if (storageFeeFrom > 0) {
                _timeStorageFeePaid[from] = block.timestamp;
                _endGracePeriod(from);
            }
        }

        // If first time receiving coins, set the grace period
        // and start the the activity clock and storage fee clock
        if (_timeStorageFeePaid[to] == 0) {
            // We may change the grace period in the future so we want to
            // preserve it per address so there is no retroactive deduction
            _storageFeeGracePeriod[to] = _storageFeeGracePeriodDays;
            _timeLastActivity[to] = block.timestamp;
            _timeStorageFeePaid[to] = block.timestamp;
        }

        // Fee transfer on `to` address
        if (storageFeeTo > 0) {
            emit Transfer(to, _feeAddress, storageFeeTo);
            _timeStorageFeePaid[to] = block.timestamp;
            _endGracePeriod(to);
        } else if (balanceToBefore < MIN_BALANCE_FOR_FEES) {
            // MIN_BALANCE_FOR_FEES is the minimum amount in which a storage fee
            // would be due after a sigle day, so if the balance is above that,
            // the storage fee would always be greater than 0.
            //
            // This avoids the following condition:
            // 1. User receives tokens
            // 2. Users sends all but a tiny amount to another address
            // 3. A year later, the user receives more tokens. Because
            // their previous balance was super small, there were no appreciable
            // storage fee, therefore the storage fee clock was not reset
            // 4. User now owes storage fees on entire balance, as if they
            // held tokens for 1 year, instead of resetting the clock to now.
            _timeStorageFeePaid[to] = block.timestamp;
        }
    }

    /**
     * @dev Apply storage fee deduction
     * @param account The account to pay storage fees
     * @return A uint256 representing the storage fee paid
     */
    function _payStorageFee(address account) internal returns (uint256) {
        uint256 storeFee = calcStorageFee(account);
        if (storeFee == 0) {
            return 0;
        }

        // Reduce account balance and add to fee address
        _balances[account] = _balances[account] - (storeFee);
        _balances[_feeAddress] = _balances[_feeAddress] + (storeFee);
        emit Transfer(account, _feeAddress, storeFee);
        _timeStorageFeePaid[account] = block.timestamp;
        _endGracePeriod(account);
        return storeFee;
    }

    /**
     * @dev Apply inactive fee deduction
     * @param account The account to pay inactive fees
     * @return A uint256 representing the inactive fee paid
     */
    function _payInactiveFee(address account) internal returns (uint256) {
        uint256 fee = _calcInactiveFee(
            _balances[account],
            daysSinceActivity(account),
            _inactiveFeePerYear[account],
            _inactiveFeePaid[account]
        );

        if (fee == 0) {
            return 0;
        }

        _balances[account] = _balances[account] - (fee);
        _balances[_feeAddress] = _balances[_feeAddress] + (fee);
        _inactiveFeePaid[account] = _inactiveFeePaid[account] + (fee);
        emit Transfer(account, _feeAddress, fee);
        return fee;
    }

    function _shouldMarkInactive(address account) internal view returns (bool) {
        // Can only mark an account as inactive if
        //
        // 1. it's not fee exempt
        // 2. it has a balance
        // 3. it's been over INACTIVE_THRESHOLD_DAYS since last activity
        // 4. it's not already marked inactive
        // 5. the storage fees owed already consume entire balance
        if (
            account != address(0) &&
            _balances[account] > 0 &&
            daysSinceActivity(account) >= INACTIVE_THRESHOLD_DAYS &&
            !isInactive(account) &&
            !isAllFeeExempt(account) &&
            (_balances[account] - calcStorageFee(account)) > 0
        ) {
            return true;
        }
        return false;
    }

    /**
     * @dev Mark an account as inactive. The function will automatically deduct
     * owed storage fees and inactive fees in one go.
     *
     * @param account The account to mark inactive
     */
    function _setInactive(address account) internal {
        // First get owed storage fees
        uint256 storeFee = calcStorageFee(account);
        uint256 snapshotBalance = _balances[account] - (storeFee);

        // all _setInactive calls are wrapped in _shouldMarkInactive, which
        // already checks this, so we shouldn't hit this condition
        assert(snapshotBalance > 0);

        // Set the account inactive on deducted balance
        _inactiveFeePerYear[account] = _calcInactiveFeePerYear(snapshotBalance);
        emit AccountInactive(account, _inactiveFeePerYear[account]);
        uint256 inactiveFees = _calcInactiveFee(
            snapshotBalance,
            daysSinceActivity(account),
            _inactiveFeePerYear[account],
            0
        );

        // Deduct owed storage and inactive fees
        uint256 fees = storeFee + (inactiveFees);
        _balances[account] = _balances[account] - (fees);
        _balances[_feeAddress] = _balances[_feeAddress] + (fees);
        _inactiveFeePaid[account] = _inactiveFeePaid[account] + (inactiveFees);
        emit Transfer(account, _feeAddress, fees);

        // Reset storage fee clock if storage fees paid
        if (storeFee > 0) {
            _timeStorageFeePaid[account] = block.timestamp;
            _endGracePeriod(account);
        }
    }

    /**
     * @dev Update the activity clock on an account thats originated a transaction.
     * If the account has previously been marked inactive or should have been
     * marked inactive, it will opportunistically collect those owed fees.
     *
     * @param account The account to update activity
     */
    function _updateActivity(address account) internal {
        // Cache has the ability to force collecting storage and inactivity fees,
        // but in the event an address was missed, can we still detect if the
        // account was inactive when they next transact
        //
        // Here we simply set the account as being inactive, collect the previous
        // storage and inactive fees that were owed, and then reactivate the account
        if (_shouldMarkInactive(account)) {
            // Call will pay existing storage fees before marking inactive
            _setInactive(account);
        }

        // Pay remaining fees and reset fee clocks
        if (isInactive(account)) {
            _payInactiveFee(account);
            _inactiveFeePerYear[account] = 0;
            _timeStorageFeePaid[account] = block.timestamp;
            emit AccountReActive(account);
        }

        // The normal case will just hit this and update
        // the activity clock for the account
        _timeLastActivity[account] = block.timestamp;
    }

    /**
     * @dev Turn off storage fee grace period for an address the first
     * time storage fees are paid (after grace period has ended)
     * @param account The account to turn off storage fee grace period
     */
    function _endGracePeriod(address account) internal {
        _storageFeeGracePeriod[account] = 0;
    }

    /**
     * @dev Simulate the transfer from one address to another see final balances and associated fees
     * @param from address The address which you want to send tokens from
     * @param to address The address which you want to transfer to
     * @return a uint256 array of 5 values representing the
     * [0] storage fees `from`
     * [1] storage fees `to`
     * [2] transfer fee `from`
     * [3] final `from` balance
     * [4] final `to` balance
     */


    /**
     * @dev Calculate the amount of inactive fees due per year on the snapshot balance.
     * Should return 50 basis points or 1 token minimum.
     *
     * @param snapshotBalance The balance of the account when marked inactive
     * @return uint256 the inactive fees due each year
     */
    function _calcInactiveFeePerYear(uint256 snapshotBalance)
    internal
    pure
    returns (uint256)
    {
        uint256 inactiveFeePerYear = (snapshotBalance * (TOKEN)) /
        (INACTIVE_FEE_DENOMINATOR);
        if (inactiveFeePerYear < TOKEN) {
            return TOKEN;
        }
        return inactiveFeePerYear;
    }

    /**
     * @dev Calculate inactive fees due on an account
     * @param balance The current account balance
     * @param daysInactive The number of days the account has been inactive
     * @param feePerYear The inactive fee per year based on snapshot balance
     * @param paidAlready The amount of inactive fees that have been paid already
     * @return uint256 for inactive fees due
     */
    function _calcInactiveFee(
        uint256 balance,
        uint256 daysInactive,
        uint256 feePerYear,
        uint256 paidAlready
    ) internal pure returns (uint256) {
        uint256 daysDue = daysInactive - (INACTIVE_THRESHOLD_DAYS);
        uint256 totalDue = (feePerYear * (TOKEN) * (daysDue)) /
        (YEAR) /
        (TOKEN) -
        (paidAlready);

        // The fee per year can be off by 0.00000001 so we can collect
        // the final dust after 200 years
        if (totalDue > balance || balance - (totalDue) <= 200) {
            return balance;
        }
        return totalDue;
    }
}