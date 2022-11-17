// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;
pragma experimental ABIEncoderV2;

contract TraceableFarms {

    // Dirección del propietario del contrato (quien lo publica)
    address public owner;

    // Constructor
    constructor() {
        owner = msg.sender;
    }

    // Estructura de consorcio de productores agrícolas
    struct Consortium {
        string id;
        string name;
    }

    // Estructura de una empresa que forma parte de un consorcio de productores agrícolas
    struct Company {
        string nif;
        string bussinessName;
        string location;
        string locationCoordinates;
        string informationalResourceUrl;
        Consortium consortium;
    }

    // Estructura de tipo de acreditación
    struct AccreditationType {
        string id;
        string name;
    }

    // Estructura de acreditación
    struct Accreditation {
        string id;
        string name;
        AccreditationType accreditationType;
    }

    // Estructura de acreditación de una empresa
    struct AccreditationCompany {
        Accreditation accreditation;
        string checker;
    }

    // Estructura de lote de un producto agrícola
    struct Batch {
        string id;
        string number;
        string date;
        string productName;
        string productVariety;
        string productDescription;
        string productPhotoUrl;
        Company company;
    }

    // Estructura de origen de un lote (de un producto agrícola)
    struct BatchOrigin {
        string description;
        string location;
        string locationCoordinates;
    }

    // Estructura de proceso de un lote (de un producto agrícola)
    struct BatchProcess {
        string name;
        string description;
    }

    // Estructura de tipo de huella
    struct FootprintType {
        string id;
        string name;
        string unitMeasurementName;
        string unitMeasurementSymbol;
    }

    // Mapping para persistencia de consorcios
    mapping (string => Consortium) private consortiums;

    // Mapping para persistencia de empresas
    mapping (string => Company) private companies;

    // Mapping para persistencia de tipos de acreditación
    mapping (string => AccreditationType) private accreditationTypes;

    // Mapping para persistencia de acreditaciones
    mapping (string => Accreditation) private accreditations;

    // Mapping para persistencia de acreditaciones de una empresa
    mapping (string => AccreditationCompany[]) private accreditationCompanies;

    // Mapping para persistencia de lotes de producto
    mapping (string => Batch) private batches;

    // Mapping para persistencia de origenes de un lote
    mapping (string => BatchOrigin[]) private batchOrigins;

    // Mapping para persistencia de procesos de un lote
    mapping (string => BatchProcess[]) private batchProcesses;

    // Mapping para persistencia de tipos de huella
    mapping (string => FootprintType) private footprintTypes;


    function setConsortium(string memory _id, string memory _name) public {
        consortiums[_id] = Consortium(_id, _name);
    }

    function getConsortium(string memory _id) public view returns (Consortium memory) {
        return consortiums[_id];
    }

    function setCompany(string memory _nif, string memory _bussinessName, string memory _location, string memory _locationCoordinates, string memory _informationalResourceUrl, string memory _consortiumId) public {
        companies[_nif] = Company(_nif, _bussinessName, _location, _locationCoordinates, _informationalResourceUrl, getConsortium(_consortiumId));
    }

    function getCompany(string memory _nif) public view returns (Company memory) {
        return companies[_nif];
    }

    function setAccreditationType(string memory _id, string memory _name) public {
        accreditationTypes[_id] = AccreditationType(_id, _name);
    }

    function getAccreditationType(string memory _id) public view returns (AccreditationType memory) {
        return accreditationTypes[_id];
    }

    function setAccreditation(string memory _id, string memory _name, string memory _accreditationTypeId) public {
        accreditations[_id] = Accreditation(_id, _name, getAccreditationType(_accreditationTypeId));
    }

    function getAccreditation(string memory _id) public view returns (Accreditation memory) {
        return accreditations[_id];
    }    

    function setAccreditationCompany(string memory _nif, string memory _accreditationId, string memory _checker) public {
        accreditationCompanies[_nif].push(
            AccreditationCompany(getAccreditation(_accreditationId), _checker)
        );
    }

    function getAccreditationCompany(string memory _nif) public view returns (AccreditationCompany[] memory) {
        return accreditationCompanies[_nif];
    }

    function setBatch(string memory _id, string memory _number, string memory _date, string memory _productName, string memory _productVariety, string memory _productDescription, string memory _productPhotoUrl, string memory _companyId) public {
        batches[_id] = Batch(_id, _number, _date, _productName, _productVariety, _productDescription, _productPhotoUrl, getCompany(_companyId));
    }

    function getBatch(string memory _id) public view returns (Batch memory) {
        return batches[_id];
    }

    function setBatchOrigin(string memory _batchId, string memory _description, string memory _location, string memory _locationCoordinates) public {
        batchOrigins[_batchId].push(
            BatchOrigin(_description, _location, _locationCoordinates)
        );
    }

    function getBatchOrigin(string memory _batchId) public view returns (BatchOrigin[] memory) {
        return batchOrigins[_batchId];
    }

    function setBatchProcess(string memory _batchId, string memory _name, string memory _description) public {
        batchProcesses[_batchId].push(
            BatchProcess(_name, _description)
        );
    }

    function getBatchProcess(string memory _batchId) public view returns (BatchProcess[] memory) {
        return batchProcesses[_batchId];
    }

    function setFootprintType(string memory _id, string memory _name, string memory _unitMeasurementName, string memory _unitMeasurementSymbol) public {
        footprintTypes[_id] = FootprintType(_id, _name, _unitMeasurementName, _unitMeasurementSymbol);
    }

    function getFootprintType(string memory _id) public view returns (FootprintType memory) {
        return footprintTypes[_id];
    }

    


}
