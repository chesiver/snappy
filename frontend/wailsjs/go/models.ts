export namespace main {
	
	export class ListDirectoryEntriesResult {
	    curDir: string;
	    fileInfos: ssh.FileInfo[];
	
	    static createFrom(source: any = {}) {
	        return new ListDirectoryEntriesResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.curDir = source["curDir"];
	        this.fileInfos = this.convertValues(source["fileInfos"], ssh.FileInfo);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace ssh {
	
	export class FileInfo {
	    name: string;
	    modTime: string;
	    isDir: boolean;
	
	    static createFrom(source: any = {}) {
	        return new FileInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.modTime = source["modTime"];
	        this.isDir = source["isDir"];
	    }
	}

}

